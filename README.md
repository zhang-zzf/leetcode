# Algorithm proj

## leetcode

### Go-GoLand

```text
# install leetcode plugin
# TemplateFilePath
/Users/admin/Workspace/github/go-leetcode

# Custom Template 
## FileName 
$!{q.frontendQuestionId}_${q.title}_test

## Template
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

${q.code}
```

### Python-PyCharm

```text
# install leetcode plugin
# TemplateFilePath
/Users/admin/Workspace/github/go-leetcode
```
