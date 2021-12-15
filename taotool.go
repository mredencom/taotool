package taotool

import (
	"github.com/mredencom/taotool/tarray"
	"github.com/mredencom/taotool/tlist"
	"github.com/mredencom/taotool/tmap"
	"github.com/mredencom/taotool/tpool"
	"github.com/mredencom/taotool/tqueue"
	"github.com/mredencom/taotool/tring"
	"github.com/mredencom/taotool/tset"
	"github.com/mredencom/taotool/ttree"
	"github.com/mredencom/taotool/ttype"
	"github.com/mredencom/taotool/tvar"
)

/**
导出工具包
*/

var (
	ArrayUtil tarray.TArray
	ListUtil  tlist.TList
	MapUtil   tmap.TMap
	PoolUtil  tpool.TPool
	QueueUtil tqueue.TQueue
	RingUtil  tring.TRing
	SetUtil   tset.TSet
	TreeUtil  ttree.TTree
	TypeUtil  ttype.TType
	VarUtil   tvar.TVar
)
