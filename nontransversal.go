import "os"
import "strings"

/*
 * Returns true if path will upward transversal (aka transversal attack).
 * Returns false if the path does not contain a upward transversal.
 *
 * For example, these will contain an upward transversal:
 *     "/.."
 *     "/www/../.."
 *     "/../etc/passwd"
 *     "/www/../../etc/passwd"
 *     "/../../../../../../../../etc/passwd"
 *
 * And these would NOT contain upward transversal:
 *     "/"
 *     "/www/../www2"
 *     "/www/../www2/file"
 *
 */
func isUpwardTransversal(path string) bool {
	parts := strings.Split(path, string(os.PathSeparator))
	var transversal int
	for _,p := range parts {
		if p == ".." {
			// they went up a directory
			transversal--
		} else if p == "." || p == "" {
			// they stayed in the same directory (no transversal)
		} else {
			// they went down a directory
			transversal++
		}
		// if they're negative, that means they went above more directories than
		// they did go down.
		if transversal < 0 {
			return true
		}
	}
	return false
}
