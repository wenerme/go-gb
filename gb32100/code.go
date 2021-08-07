package gb32100

import (
	"fmt"

	"github.com/wenerme/go-gb/gb11714"
)

// Code Unified Social Credit Identifier - 统一社会信用代码
type Code struct {
	RegDeptCode  string // 1 登记管理部门码
	OrgTypeCode  string // 1 机构类别码
	DivisionCode string // 6 登记管理机关行政区划码 GB/T 2260
	OrgCode      string // 9 主体标识码/组织机构代码 GB 11714
	Sum          string // 1 校验码 0-9,Y GB/T 17710
}

func Next(s string) (string, error) {
	c, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	next, err := c.Next()
	if err != nil {
		return "", err
	}
	return next.String(), nil
}

func Prev(s string) (string, error) {
	c, err := ParseCode(s)
	if err != nil {
		return "", err
	}
	next, err := c.Prev()
	if err != nil {
		return "", err
	}
	return next.String(), nil
}

func (u Code) RegDeptName() string {
	return _regDeptCodeName[u.RegDeptCode]
}

func (u Code) OrgTypeName() string {
	m := _orgTypeCodeName[u.RegDeptCode]
	if m == nil {
		return ""
	}
	return m[u.OrgTypeCode]
}

func (u Code) String() string {
	return fmt.Sprintf("%v%v%v%v%v", u.RegDeptCode, u.OrgTypeCode, u.DivisionCode, u.OrgCode, u.Sum)
}

func (u Code) CalcSum() string {
	s := fmt.Sprintf("%v%v%v%v", u.RegDeptCode, u.OrgTypeCode, u.DivisionCode, u.OrgCode)
	sum, _ := Sum(s)
	return sum
}

func (u Code) Next() (*Code, error) {
	neo, err := gb11714.Next(u.OrgCode)
	if err != nil {
		return nil, err
	}
	c := &Code{
		RegDeptCode:  u.RegDeptCode,
		OrgTypeCode:  u.OrgTypeCode,
		DivisionCode: u.DivisionCode,
		OrgCode:      neo,
	}
	c.Sum = c.CalcSum()
	return c, nil
}

func (u Code) Prev() (*Code, error) {
	neo, err := gb11714.Prev(u.OrgCode)
	if err != nil {
		return nil, err
	}
	c := &Code{
		RegDeptCode:  u.RegDeptCode,
		OrgTypeCode:  u.OrgTypeCode,
		DivisionCode: u.DivisionCode,
		OrgCode:      neo,
	}
	c.Sum = c.CalcSum()
	return c, nil
}

func (u Code) IsValid() bool {
	return gb11714.IsValid(u.OrgCode) && u.Sum == u.CalcSum()
}

func ParseCode(s string) (u *Code, err error) {
	if len(s) != 18 {
		return nil, fmt.Errorf("不是18位统一信用代码: %v", len(s))
	}

	u = &Code{
		RegDeptCode:  s[0:1],
		OrgTypeCode:  s[1:2],
		DivisionCode: s[2:8],
		OrgCode:      s[8:17],
		Sum:          s[17:18],
	}
	return
}

var (
	_regDeptCodeName = map[string]string{"1": "机构编制", "2": "外交", "3": "司法行政", "4": "文化", "5": "民政", "6": "旅游", "7": "宗教", "8": "工会", "9": "工商", "A": "中央军委改革和编制办公室", "N": "农业", "Y": "其他"}
	_orgTypeCodeName = map[string]map[string]string{
		"1": {"1": "机关", "2": "事业单位", "3": "编办直接管理机构编制的群众团体", "9": "其他"},
		"2": {"1": "外国常驻新闻机构", "9": "其他"},
		"3": {"1": "律师执业机构", "2": "公证处", "3": "基层法律服务所", "4": "司法鉴定机构", "5": "仲裁委员会", "9": "其他"},
		"4": {"1": "外国在华文化中心", "9": "其他"},
		"5": {"1": "社会团体", "2": "民办非企业单位", "3": "基金会", "9": "其他"},
		"6": {"1": "外国旅游部门常驻代表机构", "2": "港澳台地区旅游部门常驻内地（大陆）代表机构", "9": "其他"},
		"7": {"1": "宗教活动场所", "2": "宗教院校", "9": "其他"},
		"8": {"1": "基层工会", "9": "其他"},
		"9": {"1": "企业", "2": "个体工商户", "3": "农民专业合作社", "9": "其他"},
		"A": {"1": "军队事业单位", "9": "其他"},
		"N": {"1": "组级集体经济组织", "2": "村级集体经济组织", "3": "乡镇级集体经济组织", "9": "其他"},
		"Y": {"1": "其他"},
	}
)
