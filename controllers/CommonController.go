package controllers

const (
	SuccessMessage = "操作成功"
	ErrNotEmpty    = "不能为空"

	SuccessCode         = 0
	BindErrCode         = -1
	FilterErrCode       = -2
	InsertUpdateErrCode = -3
	DelErrCode          = -4
	GetDetailErrCode    = -5

	//上传文件类型
	UploadExcel   = 1
	UploadPicture = 2

	BaseLayoutPage = "base/layout_page.html"
	HomeBaseLayout = "home/base/layout.html"
	HomeTplPath    = "home"
	NotFundCode    = "404"
	ForbidCode     = "403"
)

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {

	if pageNo <= 0 {
		pageNo = 1
	}
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

// 通过map主键唯一的特性过滤重复元素
func FilterStrSlice(slc []string) []string {

	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
