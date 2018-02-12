/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

// Setting represents a setting
type Setting struct {
	id              int
	parentID        int
	allowUnknownArg bool
	err             error
}
