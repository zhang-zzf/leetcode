import static org.assertj.core.api.BDDAssertions.then;

import org.junit.jupiter.api.Test;


class StrongPasswordCheckerIiTest {

    final Solution solution = new Solution();

    @Test
    void givenNormal_when_thenSuccess() {
        then(solution.strongPasswordCheckerII("IloveLe3tcode!")).isTrue();
    }

    // leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public boolean strongPasswordCheckerII(String password) {
            return password.length() >= 8
                && hasLowercaseLetter(password)
                && hasUppercaseLetter(password)
                && hasDigit(password)
                && hasSpecialLetter(password)
                && noRepeatLetter(password);
        }

        private boolean noRepeatLetter(String password) {
            // return password.chars().distinct().count() == password.length();
            for (int i = 0; i < password.length() - 1; i++) {
                if (password.charAt(i) == password.charAt(i + 1)) {
                    return false;
                }
            }
            return true;
        }

        private boolean hasSpecialLetter(String password) {
            for (char c : password.toCharArray()) {
                for (char specialLetter : "!@#$%^&*()-+".toCharArray()) {
                    if (c == specialLetter) {
                        return true;
                    }
                }
            }
            return false;
        }

        private boolean hasDigit(String password) {
            for (char c : password.toCharArray()) {
                if (Character.isDigit(c)) {
                    return true;
                }
            }
            return false;
        }

        private boolean hasUppercaseLetter(String password) {
            for (char c : password.toCharArray()) {
                if (Character.isUpperCase(c)) {
                    return true;
                }
            }
            return false;
        }

        private boolean hasLowercaseLetter(String password) {
            for (char c : password.toCharArray()) {
                if (Character.isLowerCase(c)) {
                    return true;
                }
            }
            return false;
        }
    }
    // leetcode submit region end(Prohibit modification and deletion)


}