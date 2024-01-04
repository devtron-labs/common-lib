package analytics

import (
	"github.com/devtron-labs/common-lib/analytics/pprof"
	"github.com/devtron-labs/common-lib/analytics/statsViz"
	"github.com/gorilla/mux"
)

type RouterImpl struct {
	pprofRouter    pprof.PProfRouter
	statsVizRouter statsViz.StatsVizRouter
}

func (r RouterImpl) InitAnalyticsRouter(pprofSubRouter *mux.Router, statvizSubRouter *mux.Router) {
	r.pprofRouter.InitPProfRouter(pprofSubRouter)
	r.statsVizRouter.InitStatsVizRouter(statvizSubRouter)
}
