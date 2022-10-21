/**
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Example 1:
 *
 * Input: n = 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 * Example 2:
 *
 * Input: n = 0
 * Output: 0
 *
 * Example 3:
 *
 * Input: n = 1
 * Output: 0
 *
 *
 * Constraints:
 *
 * 	0 <= n <= 5 * 10^6
 *
 */

package leetcode

func countPrimes(n int) int {
	notPrime := make([]bool, n)
	count := 0

	// sieve of eratosthenes
	// don't use sqrt, because count simultaneously
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := 2; j*i < n; j++ {
				notPrime[j*i] = true
			}
		}
	}
	return count
}

/*
 * Logic is much clearer!
 *
 * func countPrimes(n int) int {
 *     primes := make([]bool, n)
 *     for i := range primes {
 *         if i == 0  || i == 1 {
 *             primes[i] = false
 *         } else {
 *             primes[i] = true
 *         }
 *     }
 *
 *     p := 2
 *     for p*p <= n {
 *         if primes[p] == true {
 *             for i := p*p; i <= n-1; i+=p {
 *                 primes[i] = false
 *             }
 *         }
 *
 *         p++
 *     }
 *
 *     primeCount := 0
 *     for i := 0; i < len(primes); i++ {
 *         if primes[i] == true {
 *             primeCount++
 *         }
 *     }
 *     return primeCount
 * }
 */
