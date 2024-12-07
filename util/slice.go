package util

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/05 10:15 am
 */

func InSliceString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func CopySliceString(src []string) []string {
	dest := make([]string, len(src))
	for i, item := range src {
		dest[i] = item
	}
	return dest
}
