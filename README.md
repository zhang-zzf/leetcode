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

# Custom Template
## FileName
$!{q.frontendQuestionId}_${q.title}_test

## Template
import unittest


class TestSolution(unittest.TestCase):

    def setUp(self) -> None:
        self.solution = Solution()

    def test_givenNormal_when_thenSuccess(self):
        self.assertTrue(True, "ShouldBeTrue")


${q.code}

if __name__ == '__main__':
    unittest.main()

```

### Java-IDEA

```text
 install leetcode plugin
# TemplateFilePath
/Users/admin/Workspace/github/go-leetcode

# Custom Template
## FileName
$!velocityTool.camelCaseName(${question.titleSlug})Test

## Template
import org.junit.jupiter.api.Test;
import static org.assertj.core.api.BDDAssertions.then;


public class $!velocityTool.camelCaseName(${question.titleSlug})Test {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
       
    }
    
    ${question.code}
    
}
```

把 `leetcode/editor/cn` 目录标记为 `Test Sources Root`
