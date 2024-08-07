package service

type HotSearchServiceInterface interface {
	// V2EX
	GetV2EX() []map[string]interface{}
	GetITHome() []map[string]interface{}
	// 知乎
	GetZhiHu() []map[string]interface{}
	// 微博
	GetWeiBo() []map[string]interface{}
	//贴吧
	GetTieBa() []map[string]interface{}
	GetDouBan() []map[string]interface{}
	GetTianYa() []map[string]interface{}
	GetHuPu() []map[string]interface{}
	GetGitHub() []map[string]interface{}
	GetBaiDu() []map[string]interface{}
	Get36Kr() []map[string]interface{}
	GetQDaily() []map[string]interface{}
	GetGuoKr() []map[string]interface{}
	GetHuXiu() []map[string]interface{}
	GetDBMovie() []map[string]interface{}
	GetZHDaily() []map[string]interface{}
	GetSegmentfault() []map[string]interface{}
	GetHacPai() []map[string]interface{}
	GetWYNews() []map[string]interface{}
	GetWaterAndWood() []map[string]interface{}
	GetNGA() []map[string]interface{}
	GetCSDN() []map[string]interface{}
	GetWeiXin() []map[string]interface{}
	GetKD() []map[string]interface{}
	GetMop() []map[string]interface{}
	GetChiphell() []map[string]interface{}
	GetJianDan() []map[string]interface{}
	GetChouTi() []map[string]interface{}
}
