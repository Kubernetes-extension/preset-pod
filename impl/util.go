package impl

import (
	"reflect"
	"regexp"
	"strings"
)

const (
	CalicoIPAddr           = "cni.projectcalico.org~1ipAddrs" // cni.projectcalico.org/ipAddrs /为特殊字符在jsonPatch中要修改为~1
	RequiredPodAnnotations = "fix.pod.ip"
)

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

// 检查RequiredPodAnnotations注解是否存在
func PodAnnotations(annotations map[string]string) (string, bool) {
	return "", true
}

// 为Pod指定特定的Node
func mutateNodeName(nodeName string) (patch []patchOperation) {
	return append(patch, patchOperation{
		Op:    "add",
		Path:  "/spec/nodeName",
		Value: nodeName,
	})
}

// 为Pod添加注解使用calico 'cni.projectcalico.org/ipAddrs' 这个特性
func addAnnotation(ipAddr string) (patch []patchOperation) {
	return append(patch, patchOperation{
		Op:    "add",
		Path:  "/metadata/annotations/" + CalicoIPAddr,
		Value: ipAddr,
	})
}

// CheckIp 检查IP地址是否合法
func CheckIp(ip string) bool {
	//addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, ip); match {
		return true
	}
	return false
}

// CheckPort 检查Port是否合法
func CheckPort(port string) bool {
	regStr := `^([1-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-5]{2}[0-3][0-5])$`
	if match, _ := regexp.MatchString(regStr, port); match {
		return true
	}
	return false
}

// CheckNotDuplicate　用于slice item 重复检查
func CheckNotDuplicate(list []string) bool {
	tmp := make(map[string]string)
	for _, i := range list {
		tmp[i] = ""
	}
	if len(tmp) != len(list) {
		return false
	}
	return true
}

// EqualSlice 比较slice是否相等
func EqualSlice(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

// GetDeploymentNameByPod 通过Pod获取Deployment Name
func GetDeploymentNameByPod(name string) string {
	nameSlice := strings.Split(name, "-")
	if len(nameSlice) > 3 {
		return strings.Join(nameSlice[:len(nameSlice)-2], "-")
	} else {
		return nameSlice[0]
	}
}
