/**
  @jeffotoni
  04/01/2019
  Front-end in Go server
*/

package assets

// GoMustAssetString returns the asset
func GoMustAssetString(strtext string) string {
	return string(MustAsset(strtext))
}
