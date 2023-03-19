package db

import "scancenter/models/scan"

//查询sdkappid信息
func (d *Connect) GetSdkAppidByAppid(req scan.GetSdkappidReq)([]BizId,error){
	var resp []BizId
	db := d.TestDataDb
	params := make([]interface{}, 0)
	//参数绑定
	sql := "select bizid from BizInfo where appid = ?"
	params = append(params,req.Appid)
	if err:=db.Raw(sql ,params...).Scan(&resp).Error;err!=nil{
		return nil,err
	}
	return resp,nil
}