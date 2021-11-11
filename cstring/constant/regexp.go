package constant

// Value
//  Description: 值
//  Author: Kevin·CC
//  return string
func (m MyRegexp) Value() string {
	return string(m)
}

type MyRegexp string

const (
	General            MyRegexp = "^\\w+$"
	Numbers            MyRegexp = "\\d+"
	Word               MyRegexp = "[a-zA-Z]+"
	Chinese            MyRegexp = "[\u4E00-\u9FFF]"
	Chineses           MyRegexp = Chinese + "+"
	GroupVar           MyRegexp = "\\$(\\d+)"
	Ipv4               MyRegexp = "^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)$"
	Ipv6               MyRegexp = "(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]+|::(ffff(:0{1,4})?:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))"
	Money              MyRegexp = "^(\\d+(?:\\.\\d+)?)$"
	Email              MyRegexp = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"
	Mobile             MyRegexp = "(?:0|86|\\+86)?1[3-9]\\d{9}"
	MobileHk           MyRegexp = "(?:0|852|\\+852)?\\d{8}"
	MobileTw           MyRegexp = "(?:0|886|\\+886)?(?:|-)09\\d{8}"
	MobileMo           MyRegexp = "(?:0|853|\\+853)?(?:|-)6\\d{7}"
	Tel                MyRegexp = "(010|02\\d|0[3-9]\\d{2})-?(\\d{6,8})"
	Tel400_800         MyRegexp = "0\\d{2,3}[\\- ]?[1-9]\\d{6,7}|[48]00[\\- ]?[1-9]\\d{6}"
	CitizenId          MyRegexp = "[1-9]\\d{5}[1-2]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3}(\\d|X|x)"
	ZipCode            MyRegexp = "^(0[1-7]|1[0-356]|2[0-7]|3[0-6]|4[0-7]|5[0-7]|6[0-7]|7[0-5]|8[0-9]|9[0-8])\\d{4}|99907[78]$"
	Birthday           MyRegexp = "^(\\d{2,4})([/\\-.年]?)(\\d{1,2})([/\\-.月]?)(\\d{1,2})日?$"
	Uri                MyRegexp = "^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?"
	Url                MyRegexp = "[a-zA-Z]+://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"
	UrlHttp            MyRegexp = "(https?|ftp|file)://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"
	GeneralWithChinese MyRegexp = "^[\u4E00-\u9FFF\\w]+$"
	Uuid               MyRegexp = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UuidSimple         MyRegexp = "^[0-9a-fA-F]{32}$"
	MacAddress         MyRegexp = "((?:[a-fA-F0-9]{1,2}[:-]){5}[a-fA-F0-9]{1,2})|0x(\\d{12}).+ETHER"
	Hex                MyRegexp = "^[a-fA-F0-9]+$"
	Time               MyRegexp = "\\d{1,2}:\\d{1,2}(:\\d{1,2})?"
	PlateNumber        MyRegexp = "^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[ABCDEFGHJK])|([ABCDEFGHJK]([A-HJ-NP-Z0-9])[0-9]{4})))|" +
		//https://gitee.com/loolly/hutool/issues/I1BJHE?from=project-issue
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领]\\d{3}\\d{1,3}[领])|" +
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$"
	CreditCode        MyRegexp = "^[0-9A-HJ-NPQRTUWXY]{2}\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$"
	CarVin            MyRegexp = "^[A-Za-z0-9]{17}$"
	CarDrivingLicence MyRegexp = "^[0-9]{12}$"
)
