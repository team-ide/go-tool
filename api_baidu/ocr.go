package api_baidu

import (
	"bytes"
	"encoding/base64"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"image"
	_ "image/jpeg"
	"image/png"
	"net/url"
	"os"
)

var OcrAccurateBasicUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

type OcrAccurateBasicResponse struct {
	LogId          uint64   `json:"log_id"`           // 唯一的log id，用于问题定位
	Direction      int32    `json:"direction"`        // 图像方向，当 detect_direction=true 时返回该字段。 - 1：未定义， 0：正向， 1：逆时针90度， 2：逆时针180度， 3：逆时针270度
	WordsResultNum uint32   `json:"words_result_num"` // 识别结果数，表示words_result的元素个数
	WordsResult    []*Words `json:"words_result"`     //识别结果数组
}

type Words struct {
	Words string `json:"words"` // 识别结果字符串
}

func GetImage64ByFile(filePath string) (base64Str string, err error) {
	fSrc, err := os.Open(filePath)
	if err != nil {
		util.Logger.Error("os.Open error", zap.Error(err))
		return
	}
	defer func() { _ = fSrc.Close() }()
	img, _, err := image.Decode(fSrc)
	if err != nil {
		util.Logger.Error("image.Decode", zap.Error(err))
		return
	}
	// 这里的resImg是一个 image.Image 类型的变量
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		util.Logger.Error("png.Encode error", zap.Error(err))
		return
	}
	// 将字节切片转换为Base64字符串
	base64Str = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

func OcrAccurateBasicImageByFile(accessToken string, filePath string) (res *OcrAccurateBasicResponse, err error) {
	base64Str, err := GetImage64ByFile(filePath)
	if err != nil {
		return
	}
	res, err = OcrAccurateBasicImage(accessToken, base64Str)
	if err != nil {
		return
	}
	return
}

func OcrAccurateBasicImage(accessToken string, imageBase64 string) (res *OcrAccurateBasicResponse, err error) {

	// 创建一个表单
	form := url.Values{}
	// 是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度
	form.Add("detect_direction", "true")
	// 是否开启行级别的多方向文字识别
	form.Add("multidirectional_recognize", "true")
	form.Add("image", imageBase64)

	apiUrl := OcrAccurateBasicUrl + "?access_token=" + accessToken
	res, err = util.PostForm(apiUrl, form, &OcrAccurateBasicResponse{})
	if err != nil {
		return
	}

	return
}
