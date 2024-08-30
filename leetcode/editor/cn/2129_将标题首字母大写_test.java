//给你一个字符串 title ，它由单个空格连接一个或多个单词组成，每个单词都只包含英文字母。请你按以下规则将每个单词的首字母 大写 ： 
//
// 
// 如果单词的长度为 1 或者 2 ，所有字母变成小写。 
// 否则，将单词首字母大写，剩余字母变成小写。 
// 
//
// 请你返回 大写后 的 title 。 
//
// 
//
// 示例 1： 
//
// 输入：title = "capiTalIze tHe titLe"
//输出："Capitalize The Title"
//解释：
//由于所有单词的长度都至少为 3 ，将每个单词首字母大写，剩余字母变为小写。
// 
//
// 示例 2： 
//
// 输入：title = "First leTTeR of EACH Word"
//输出："First Letter of Each Word"
//解释：
//单词 "of" 长度为 2 ，所以它保持完全小写。
//其他单词长度都至少为 3 ，所以其他单词首字母大写，剩余字母小写。
// 
//
// 示例 3： 
//
// 输入：title = "i lOve leetcode"
//输出："i Love Leetcode"
//解释：
//单词 "i" 长度为 1 ，所以它保留小写。
//其他单词长度都至少为 3 ，所以其他单词首字母大写，剩余字母小写。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= title.length <= 100 
// title 由单个空格隔开的单词组成，且不含有任何前导或后缀空格。 
// 每个单词由大写和小写英文字母组成，且都是 非空 的。 
// 
//
// Related Topics 字符串 
// 👍 52 👎 0


import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class CapitalizeTheTitleTest {

  final Solution solution = new Solution();

  @Test
  void givenNormal_when_thenSuccess() {
    String ans = solution.capitalizeTitle("i lOve leetcode");
    then(ans).isEqualTo("i Love Leetcode");
  }

  //leetcode submit region begin(Prohibit modification and deletion)
  class Solution {
    public String capitalizeTitle(String title) {
      StringBuilder buf = new StringBuilder();
      // mark end
      title += " ";
      int lng = title.length();
      for (int left = -1, right = 0; right < lng; right++) {
        if (title.charAt(right) != ' ') {
          if (left == -1) {// word start
            left = right;
          }
          continue;
        }
        if (right - left > 2) {
          buf.append(toUpper(title.charAt(left++)));
        }
        while (left < right) {
          buf.append(toLower(title.charAt(left++)));
        }
        buf.append(' ');
        left = -1;// remark next word start index
      }
      if (buf.length() > 0) {
        // remove end mark
        return buf.substring(0, buf.length() - 1);
      }
      return "";
    }

    private char toLower(char c) {
      if (c >= 'a') {
        return c;
      }
      return (char) (c + ('a' - 'A'));
    }

    private char toUpper(char c) {
      if (c < 'a') {
        return c;
      }
      return (char) (c + ('A' - 'a'));
    }
  }
//leetcode submit region end(Prohibit modification and deletion)


}