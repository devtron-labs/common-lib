/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package nats_lib

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	CI_RUNNER_STREAM                  string = "CI-RUNNER"
	ORCHESTRATOR_STREAM               string = "ORCHESTRATOR"
	KUBEWATCH_STREAM                  string = "KUBEWATCH"
	GIT_SENSOR_STREAM                 string = "GIT-SENSOR"
	BULK_APPSTORE_DEPLOY_TOPIC        string = "APP-STORE.BULK-DEPLOY"
	BULK_APPSTORE_DEPLOY_GROUP        string = "APP-STORE-BULK-DEPLOY-GROUP-1"
	BULK_APPSTORE_DEPLOY_DURABLE      string = "APP-STORE-BULK-DEPLOY-DURABLE-1"
	CD_STAGE_COMPLETE_TOPIC           string = "CD-STAGE-COMPLETE"
	CD_COMPLETE_GROUP                 string = "CD-COMPLETE_GROUP-1"
	CD_COMPLETE_DURABLE               string = "CD-COMPLETE_DURABLE-1"
	BULK_DEPLOY_TOPIC                 string = "CD.BULK"
	BULK_HIBERNATE_TOPIC              string = "CD.BULK-HIBERNATE"
	BULK_DEPLOY_GROUP                 string = "CD.BULK.GROUP-1"
	BULK_HIBERNATE_GROUP              string = "CD.BULK-HIBERNATE.GROUP-1"
	BULK_DEPLOY_DURABLE               string = "CD-BULK-DURABLE-1"
	BULK_HIBERNATE_DURABLE            string = "CD-BULK-HIBERNATE-DURABLE-1"
	CI_COMPLETE_TOPIC                 string = "CI-COMPLETE"
	CI_COMPLETE_GROUP                 string = "CI-COMPLETE_GROUP-1"
	CI_COMPLETE_DURABLE               string = "CI-COMPLETE_DURABLE-1"
	APPLICATION_STATUS_UPDATE_TOPIC   string = "APPLICATION_STATUS_UPDATE"
	APPLICATION_STATUS_UPDATE_GROUP   string = "APPLICATION_STATUS_UPDATE_GROUP-1"
	APPLICATION_STATUS_UPDATE_DURABLE string = "APPLICATION_STATUS_UPDATE_DURABLE-1"
	CRON_EVENTS                       string = "CRON_EVENTS"
	CRON_EVENTS_GROUP                 string = "CRON_EVENTS_GROUP-2"
	CRON_EVENTS_DURABLE               string = "CRON_EVENTS_DURABLE-2"
	WORKFLOW_STATUS_UPDATE_TOPIC      string = "WORKFLOW_STATUS_UPDATE"
	WORKFLOW_STATUS_UPDATE_GROUP      string = "WORKFLOW_STATUS_UPDATE_GROUP-1"
	WORKFLOW_STATUS_UPDATE_DURABLE    string = "WORKFLOW_STATUS_UPDATE_DURABLE-1"
	CD_WORKFLOW_STATUS_UPDATE         string = "CD_WORKFLOW_STATUS_UPDATE"
	CD_WORKFLOW_STATUS_UPDATE_GROUP   string = "CD_WORKFLOW_STATUS_UPDATE_GROUP-1"
	CD_WORKFLOW_STATUS_UPDATE_DURABLE string = "CD_WORKFLOW_STATUS_UPDATE_DURABLE-1"
	NEW_CI_MATERIAL_TOPIC             string = "NEW-CI-MATERIAL"
	NEW_CI_MATERIAL_TOPIC_GROUP       string = "NEW-CI-MATERIAL_GROUP-1"
	NEW_CI_MATERIAL_TOPIC_DURABLE     string = "NEW-CI-MATERIAL_DURABLE-1"
	CD_SUCCESS                        string = "CD.TRIGGER"
	WEBHOOK_EVENT_TOPIC               string = "WEBHOOK_EVENT"
	MSG_MAX_AGE                       int    = 86400
)

var ORCHESTRATOR_SUBJECTS = []string{BULK_APPSTORE_DEPLOY_TOPIC, BULK_DEPLOY_TOPIC, BULK_HIBERNATE_TOPIC, CD_SUCCESS, WEBHOOK_EVENT_TOPIC}
var ORCHESTRATOR_CONSUMERS = []string{BULK_APPSTORE_DEPLOY_DURABLE, BULK_DEPLOY_DURABLE, BULK_HIBERNATE_DURABLE}
var CI_RUNNER_SUBJECTS = []string{CI_COMPLETE_TOPIC, CD_STAGE_COMPLETE_TOPIC}
var CI_RUNNER_CONSUMERS = []string{CI_COMPLETE_DURABLE, CD_COMPLETE_DURABLE}
var KUBEWATCH_SUBJECTS = []string{APPLICATION_STATUS_UPDATE_TOPIC, CRON_EVENTS, WORKFLOW_STATUS_UPDATE_TOPIC, CD_WORKFLOW_STATUS_UPDATE}
var KUBEWATCH_CONSUMERS = []string{APPLICATION_STATUS_UPDATE_DURABLE, CRON_EVENTS_DURABLE, WORKFLOW_STATUS_UPDATE_DURABLE, CD_WORKFLOW_STATUS_UPDATE_DURABLE}
var GIT_SENSOR_SUBJECTS = []string{NEW_CI_MATERIAL_TOPIC}
var GIT_SENSOR_CONSUMERS = []string{NEW_CI_MATERIAL_TOPIC_DURABLE}

func GetStreamSubjects(streamName string) []string {
	var subjArr []string
	switch streamName {
	case ORCHESTRATOR_STREAM:
		subjArr = ORCHESTRATOR_SUBJECTS
	case CI_RUNNER_STREAM:
		subjArr = CI_RUNNER_SUBJECTS
	case KUBEWATCH_STREAM:
		subjArr = KUBEWATCH_SUBJECTS
	case GIT_SENSOR_STREAM:
		subjArr = GIT_SENSOR_SUBJECTS
	default:
		subjArr = []string{"hello.world"}
	}
	return subjArr
}

func GetStreamConsumers(streamName string) []string {
	var consArr []string
	switch streamName {
	case ORCHESTRATOR_STREAM:
		consArr = ORCHESTRATOR_CONSUMERS
	case CI_RUNNER_STREAM:
		consArr = CI_RUNNER_CONSUMERS
	case KUBEWATCH_STREAM:
		consArr = KUBEWATCH_CONSUMERS
	case GIT_SENSOR_STREAM:
		consArr = GIT_SENSOR_CONSUMERS
	default:
		consArr = []string{WORKFLOW_STATUS_UPDATE_DURABLE}
	}
	return consArr
}

func AddStream(js nats.JetStreamContext, streamConfig *nats.StreamConfig, streamNames ...string) error {
	var err error
	for _, streamName := range streamNames {
		streamInfo, err := js.StreamInfo(streamName)
		if err == nats.ErrStreamNotFound || streamInfo == nil {
			log.Print("No stream was created already. Need to create one.", "Stream name", streamName)
			//Stream doesn't already exist. Create a new stream from jetStreamContext
			cfgToSet := getNewConfig(streamName, streamConfig)
			_, err = js.AddStream(cfgToSet)
			if err != nil {
				log.Fatal("Error while creating stream", "stream name", streamName, "error", err)
				return err
			}
		} else if err != nil {
			log.Fatal("Error while getting stream info", "stream name", streamName, "error", err)
		} else {
			config := streamInfo.Config
			if checkConfigChangeReqd(&config, streamConfig) {
				streamConfig.Name = streamName
				streamConfig.Subjects = GetStreamSubjects(streamName)
				_, err := js.UpdateStream(streamConfig)
				if err != nil {
					log.Println("error occurred while updating stream config", "streamName", streamName, "streamConfig", streamConfig, "error", err)
				} else {
					log.Println("stream config updated successfully", "config", config, "new", streamConfig)
				}
			}
		}
		//consumers := GetStreamConsumers(streamName)
		//for _, consumer := range consumers {
		//	consumerInfo, err := js.ConsumerInfo(streamName, consumer)
		//	if err == nats.ErrConsumerNotFound && consumerInfo == nil {
		//		_, err := js.AddConsumer(streamName, &nats.ConsumerConfig{
		//			Durable:   consumer,
		//			AckPolicy: nats.AckExplicitPolicy,
		//		})
		//		if err != nil {
		//			//TODO handle error case
		//		}
		//	}
		//}
	}
	return err
}

func checkConfigChangeReqd(existingConfig *nats.StreamConfig, toUpdateConfig *nats.StreamConfig) bool {
	configChanged := false
	if toUpdateConfig.MaxAge != time.Duration(0) && toUpdateConfig.MaxAge != existingConfig.MaxAge {
		configChanged = true
	} else {
		toUpdateConfig.MaxAge = existingConfig.MaxAge
	}

	//if toUpdateConfig.Replicas != 0 && toUpdateConfig.Replicas != existingConfig.Replicas {
	//	configChanged = true
	//} else {
	//	toUpdateConfig.Replicas = existingConfig.Replicas
	//} commented as retention policy cannot be updated

	//if toUpdateConfig.Retention != existingConfig.Retention {
	//	configChanged = true
	//} else {
	//	toUpdateConfig.Retention = existingConfig.Retention
	//} commented as retention policy cannot be updated
	return configChanged
}

func getNewConfig(streamName string, toUpdateConfig *nats.StreamConfig) *nats.StreamConfig {
	cfg := &nats.StreamConfig{
		Name:     streamName,
		Subjects: GetStreamSubjects(streamName),
	}

	if toUpdateConfig.MaxAge != time.Duration(0) {
		cfg.MaxAge = toUpdateConfig.MaxAge
	}
	if toUpdateConfig.Replicas > 0 {
		cfg.Replicas = toUpdateConfig.Replicas
	}
	if toUpdateConfig.Retention != nats.RetentionPolicy(0) {
		cfg.Retention = toUpdateConfig.Retention
	}
	cfg.Retention = nats.RetentionPolicy(2)
	return cfg
}

func getMaxAge() time.Duration {
	natsMaxAgeStr := os.Getenv("NATS_STREAM_MAX_AGE")
	msgMaxAge, err := strconv.Atoi(natsMaxAgeStr)
	if err != nil {
		log.Println("error occurred while converting maxAge to integer", "natsMaxAgeStr", natsMaxAgeStr, "error", err)
		msgMaxAge = MSG_MAX_AGE
	}
	return time.Duration(msgMaxAge) * time.Second
}
