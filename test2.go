package main

import (
	"fmt"
	"git.ezbuy.me/ezbuy/base/misc/context"
	"git.ezbuy.me/ezbuy/spike/rpc/spike"
)

func main() {
	arr := []int {1,2,3,4,5,6,7,8,9,10}
	/*as := T(arr,3)
	if as != nil {
		for _,a := range as {
			i,ok :=  a.(int)
			fmt.Println("--------------",ok,i)
		}
	}*/
	for i := 0; i < len(arr)/l+1; i++ {
		if err != nil {
			break
		}
		s := i * l
		e := 0
		if len(req.GetReq()) < (i+1)*l {
			e = len(req.GetReq())
		} else {
			e = (i + 1) * l
		}
		if s == e {
			continue
		}

		newReq := req
		newReq.Req = req.GetReq()[s:e]

		log.Debugf("GetOracleSimpleInfo %v, %d,%d", newReq.GetReq(), s, e)

		wg.Add(1)
		go func(c context.T, r spike.OracleGetMultiProduct) {
			res, e := spike.GetOracle().GetSimpleInfo(ctx.GoDump(), newReq)
			if e == nil && res != nil {
				mu.Lock()
				respProducts = append(respProducts, res.GetResult()...)
				mu.Unlock()
			} else if e != nil {
				err = e
				log.Error(e)
			}
			wg.Done()
		}(ctx.GoDump(), *newReq)
	}
}


/*func T(arr interface{},l int) []interface{} {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return nil
	}


	as,ok := arr.([]interface{})
	fmt.Println(ok,as)

	rs := make([]interface{},0)

	t := len(as) / l
	for i:=0;i<t+1;i++ {
		s := i*l
		e := 0
		if len(as) < (i+1)*l {
			e = len(as)
		} else {
			e = (i+1)*l
		}

		if s == e {
			continue
		}
		a := as[s:e]

		fmt.Println(a,s,e)

		rs = append(rs,a...)
	}

	return rs
}*/