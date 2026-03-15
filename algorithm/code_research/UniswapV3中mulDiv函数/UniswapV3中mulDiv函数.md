# 代码探索之mulDiv
 这是DeFi 领域（特别是 Uniswap V3 FullMath 库）中最经典的 512 位高精度乘法与除法算法 (mulDiv)
 
 <details>
    <summary>Uniswap V3 FullMath 库中的mulDiv函数</summary>

```solidity
function mulDiv(
    uint256 a,
    uint256 b,
    uint256 denominator
) internal pure returns (uint256 result) {
    uint256 prod0; // Least significant 256 bits of the product
    uint256 prod1; // Most significant 256 bits of the product
    assembly {
        let mm := mulmod(a, b, not(0))
        prod0 := mul(a, b)
        prod1 := sub(sub(mm, prod0), lt(mm, prod0))
    }
    if (prod1 == 0) {
        require(denominator > 0);
        assembly {
            result := div(prod0, denominator)
        }
        return result;
    }
    require(denominator > prod1);
    uint256 remainder;
    assembly {
        remainder := mulmod(a, b, denominator)
    }
    assembly {
        prod1 := sub(prod1, gt(remainder, prod0))
        prod0 := sub(prod0, remainder)
    }
    uint256 twos = -denominator & denominator;
    assembly {
        denominator := div(denominator, twos)
    }
    assembly {
        prod0 := div(prod0, twos)
    }
    assembly {
        twos := add(div(sub(0, twos), twos), 1)
    }
    prod0 |= prod1 * twos;
    uint256 inv = (3 * denominator) ^ 2; // inverse mod 2**4
    inv *= 2 - denominator * inv; // inverse mod 2**8
    inv *= 2 - denominator * inv; // inverse mod 2**16
    inv *= 2 - denominator * inv; // inverse mod 2**32
    inv *= 2 - denominator * inv; // inverse mod 2**64
    inv *= 2 - denominator * inv; // inverse mod 2**128
    inv *= 2 - denominator * inv; // inverse mod 2**256
    result = prod0 * inv;
    return result;
}
```

</details>

## 代码解读
- 这段代码是 DeFi 领域（特别是 Uniswap V3 FullMath 库）中最经典的 512 位高精度乘法与除法算法 (mulDiv)。
- 它的核心作用是：在不丢失精度、且防止中间结果溢出的情况下，计算 (a * b) / denominator。
- 时间复杂度为O(1)

### 整体数学逻辑
- 512位乘积拆分：利用 EVM 汇编指令 mulmod 和 mul，将 a * b 的真实 512 位结果拆分为高 256 位 (prod1) 和低 256 位 (prod0)。
- 边界与溢出拦截：如果 prod1 == 0，说明没溢出，直接返回普通除法。如果 prod1 >= denominator，说明最终结果一定会大于$
2^{256} - 1$（即使除完以后也会溢出 uint256 返回值），强制拦截。
- 余数剥离：求出 (a * b) % denominator 的余数，并从 512 位数字中减去该余数，使得分子能被 denominator 绝对整除。
- 因式分解与位移：提取 denominator 的最大 2 的幂次因子 (twos)，将分子和分母同时除以 twos。这确保了新的 denominator 是一个奇数。
- 牛顿-拉弗森迭代求模逆元 (Newton-Raphson)：由于任何奇数与 2<sup>256</sup>都互质，因此该奇数在模 2<sup>256</sup>下必定存在乘法逆元。通过 6 次迭代求出逆元 inv。
- 除法转乘法：最终结果 = prod0 * inv。在模 2<sup>256</sup>的运算中，乘以逆元等价于除以原数。

### 逐段解读代码

#### 拦截高位与低位 (512-bit Representation)
```solidity
let mm := mulmod(a, b, not(0)) // a * b % (2^256 - 1)
prod0 := mul(a, b)             // a * b (低 256 位，溢出部分被截断)
prod1 := sub(sub(mm, prod0), lt(mm, prod0)) // 巧妙推导出高 256 位
```
EVM 没有 512 位原生操作码。这里利用了模运算的同余定理。`not(0)` 就是 `type(uint256).max`。通过 `mulmod` 获取真实的模值，再通过与截断后的低位 `prod0` 比较，利用高低位差值精准反推出被截断的高 256 位 `prod1`。
<details>
<summary>数学推导</summary>

##### 第一步：获取低 256 位 (`prod0`)
在 EVM 中，直接执行 `mul(a, b)` 时，如果结果超过 256 位，高位会被直接截断（丢弃），只保留低 256 位。
数学表达为：`prod0 = (a * b) % 2^256`。
因此，`prod0 := mul(a, b)` 直接获取了绝对准确的低 256 位。

##### 第二步：获取真实的模值 (`mm`)
EVM 提供了一个特殊的操作码 `mulmod(x, y, k)`，它会在底层 C++ 引擎中以 512 位精度计算 `(x * y) % k`，**中间过程绝不会溢出**。
代码中传入的 `k` 是 `not(0)`，即二进制全 1，等于 $2^{256} - 1$。
因此，`mm = (a * b) % (2^256 - 1)`。这是绝对准确的。

##### 第三步：建立高低位与 `mm` 的同余方程
基于前文的定义，真实乘积可以表示为：
`a * b = prod1 * 2^256 + prod0`

现在，我们将等式两边同时对 $(2^{256} - 1)$ 取模：
1. 左边 `(a * b) % (2^256 - 1)` 就是我们刚才算出的 `mm`。
2. 右边重点看 $2^{256}$ 这个数字。在模 $(2^{256} - 1)$ 的世界里，$2^{256}$ 等价于 1。（因为 $2^{256} - (2^{256} - 1) = 1$）。
3. 所以，右边化简为：`(prod1 * 1 + prod0) % (2^{256} - 1)`。

推导出的核心同余方程为：
**`mm ≡ prod1 + prod0 (mod 2^256 - 1)`**

##### 第四步：求解高位 `prod1`
根据同余方程，`prod1 + prod0` 只有两种情况（因为 `prod1` 和 `prod0` 都必定小于 $2^{256}$，它们的和不可能达到模数的 2 倍）：

*   **情况 A (未发生进位绕回)**：`prod1 + prod0 = mm`
    推导出 -> `prod1 = mm - prod0`。此时 `mm >= prod0`。
*   **情况 B (发生进位绕回)**：`prod1 + prod0 = mm + (2^256 - 1)`
    推导出 -> `prod1 = mm - prod0 + (2^256 - 1)`。此时因为加上了模数，必定有 `mm < prod0`。

</details>

#### 余数平滑与除法降维
```solidity
assembly {
    remainder := mulmod(a, b, denominator)
    prod1 := sub(prod1, gt(remainder, prod0)) // 借位减法
    prod0 := sub(prod0, remainder)
}
```
*  为了让后续的除法没有小数，先算出真实余数，从 `[prod1, prod0]` 这个 512 位大整数中整体减去。如果低位 `prod0` 不够减，`gt(remainder, prod0)` 返回 1，向高位 `prod1` 借位。

#### 提取 2 的幂次 (Bitwise Magic)
```solidity
uint256 twos = -denominator & denominator;
```
*   极其经典的补码位运算（Two's complement trick）。在底层二进制中，`-x & x` 会绝对精确地提取出 `x` 的最低位 `1`（即最大能整除该数的 2 的幂）。
*   随后的 `twos := add(div(sub(0, twos), twos), 1)` 是在计算 $2^{256} / twos$ 的等价汇编写法，目的是将高 256 位的有效数据向右位移，压入低 256 位 `prod0` 中。

#### 牛顿迭代求解模逆元 (Inverse Modulo $2^{256}$)
```solidity
uint256 inv = (3 * denominator) ^ 2; // Magic number 初始化
inv *= 2 - denominator * inv; // inverse mod 2**8
// ... 翻倍精度直到 2**256
result = prod0 * inv;
```
*   这是密码学中经典的求逆元算法。一旦分母提去了所有的 2（变成奇数），它在模 $2^{256}$ 域内必定有唯一逆元。每次迭代 `inv *= 2 - denominator * inv` 都会使精确度翻倍（从模 $2^4$ -> $2^8$ -> $2^{16}$ 直到 $2^{256}$）。
*   最后，利用**“除以一个数等于乘以它的倒数（逆元）”**的原理，用微秒级的乘法 `prod0 * inv` 替代了复杂的 512 位除法状态机。

## 相关技术原理
虽然代码只有短短五十行，但是涉及的数学和技术原理有七个之多，相关的数学定义以及证明过程就不在此展开了，可自行搜索。

### 同余定理 - Congruence Theorem / Modular Arithmetic
- 同余：给定一个正整数 m（模数）。如果两个整数 a 和 b 满足它们除以 m 的余数相同，那么称 a 和 b 对模 m 同余。
- 同余关系在加法、减法和乘法下是完全保持封闭和自洽的

#### 在代码中的作用
`状态空间压缩`与`防溢出降维`:每一步加法或乘法运算后立刻取模，其最终结果与“先算出巨大的真实结果，再整体取模”是绝对等价的

- 对于uint64类型的a, b, m, 通过`(a * b) mod m = ((a mod m) * (b mod m)) mod m`来防止a * b的uint64类型溢出
- 在密码学中计算 base<sup>exp</sup>(mod m) 时，利用同余幂定律，可以通过平方法将 O(N) 的连乘时间复杂度降维至 O(logN) 

#### 风险点
- 同余定理不支持直接除法
- 在Go,Java,C++中，%的本质是取余，保留被除数的符号（比如在Go中 `-5 % 3 = -2`），所以需要特别注意在计算取余前加上m

### 位运算获取最低位1的位置
在底层二进制中，`-x & x` 会绝对精确地提取出 x 的最低位 1（即最大能整除该数的 2 的幂）

### CRT - 中国剩余定理 (Chinese Remainder Theorem)
- 如果存在一组两两互质的模数，那么对于任意给定的余数，该同余线性方程必定有解；并且在模数乘积范围内，解唯一

#### 在代码中的作用
- `数据的分片 (Map)`：一个 2048 位的超大操作数 X ，计算成本极其高昂。但利用 CRT 原理，我们可以把它模上几个小质数，拆解成几个 256 位的小余数（a1,a2,...）。
- `降维计算 (Process)`：在这 几个独立的 256 位空间里分别进行计算。此时由于位数变小，CPU 的乘除法指令周期呈指数级下降，且这些计算是完全正交的，可以扔进多个 Goroutine 中并行执行。
- 结果缝合 (Reduce)：利用 CRT 构造性证明中的公式 x=∑ai*Mi*yi(mod M) ，将这些小空间里的计算结果，瞬间还原成 2048 位空间内的绝对正确答案。

#### 风险点
- 输入的两个模数必须互质：若面临非互质的工程场景，必须降级使用扩展中国剩余定理 (Generalized CRT)。此时不再使用乘积 M，而是使用最小公倍数 LCM，且方程组有解的严苛前提变为：ai ≡ aj(mod gcd(mi,mj))
- 累加器的内存边界
    - Go中写时必须全程保持在 math/big 的堆内存运算中。
    - Solidity中写时，在每一轮循环累加后，立即执行模加法，确保内存中的数据体积永远不超过 M。

### 乘法逆元 - Modular Multiplicative Inverse
离散数学和密码学（模算术）中的“倒数”。
- 在连续的实数域中，除以一个数等于乘以它的倒数。
- 在计算机底层的有限整数域（如 EVM 的 2<sup>256</sup>空间或椭圆曲线密码学中），不存在小数。为了在纯整数空间内实现“除法”，我们需要找到一个整数 inv，使得原数乘以 inv 溢出并截断后的结果刚好等于 1。这个 inv 就是原数的乘法逆元。

#### 在代码中的作用
将计算机底层极其昂贵且易引发精度丢失的“除法”指令，完美等效替换为极速的“乘法”指令。

#### 魔法初始值
对于任意奇数 d，`(3 * d) XOR 2` 的结果，必定是 d在模 16（即 2<sup>4</sup>）下的准确乘法逆元。

#### 风险点
- 乘法逆元并非永远存在。裴蜀定理规定：只有当 a 和模数 m 的最大公约数为 1，即 `gcd(a,m)=1`（两者互质）时，a 在模 m 下的逆元才存在，且唯一。
- 在 EVM 中 (m=2<sup>256</sup> )的质因数只有 2。因此，只有奇数才与 2<sup></sup>互质。这也是为什么在 Uniswap 的 mulDiv 源码中，必须先用按位与提取出分母的所有 2 的幂次因子，将分母强行变成奇数后，才能安全进入牛顿迭代求解逆元。

### 牛顿-拉弗森迭代 (Newton-Raphson Iteration) 

**迭代公式**：
$$x_{n+1} = x_n - \frac{f(x_n)}{f'(x_n)}$$

**求乘法逆元**:
*   构建反向函数：**$f(x) = \frac{1}{x} - d = 0$**
*   对其求导：$f'(x) = -\frac{1}{x^2}$
*   代入牛顿迭代公式：$x_{n+1} = x_n - \frac{\frac{1}{x_n} - d}{-\frac{1}{x^2_n}} = x_n + x_n^2 \left(\frac{1}{x_n} - d\right)$
*   化简得到最终形态：**$x_{n+1} = x_n(2 - d \cdot x_n)$**
*   **极致优雅**：这个公式里只有乘法和减法，完全消除了除法！

### 亨泽尔引理与精度翻倍 (Hensel's Lemma & Quadratic Convergence)
在模算术中，牛顿迭代表现为“精度翻倍”。
假设当前迭代值 $x_n$ 在模 $2^k$ 下是正确的逆元，即 $x_n \cdot d \equiv 1 \pmod{2^k}$。
我们可以将其写为：$x_n \cdot d = 1 + c \cdot 2^k$（$c$ 为某个常数）。
我们将 $x_{n+1}$ 乘以 $d$ 看看会发生什么：
$x_{n+1} \cdot d = x_n(2 - d \cdot x_n) \cdot d = 2(x_n \cdot d) - (x_n \cdot d)^2$
代入刚才的式子：
$= 2(1 + c \cdot 2^k) - (1 + c \cdot 2^k)^2$
$= 2 + 2c \cdot 2^k - (1 + 2c \cdot 2^k + c^2 \cdot 2^{2k})$
$= 1 - c^2 \cdot 2^{2k}$
得出结论：**$x_{n+1} \cdot d \equiv 1 \pmod{2^{2k}}$**。
每一次迭代，正确的有效位数直接翻倍（从 $k$ 位变成 $2k$ 位）。

#### 代码中
有了魔法初始值`(3 * d) XOR 2`这个模 $2^4$ 的绝对正确起点，我们只需要迭代：
*   第1次：模 $2^8$
*   第2次：模 $2^{16}$
*   第3次：模 $2^{32}$
*   第4次：模 $2^{64}$
*   第5次：模 $2^{128}$
*   第6次：模 $2^{256}$
刚好 6 行代码，将精度拉满到 256 位。

### 扩展欧几里得算法
扩展欧几里得算法（Extended Euclidean Algorithm, ExGCD）是标准欧几里得算法（辗转相除法）的进阶版。

它不仅用于求解两个整数 $a$ 和 $b$ 的最大公约数 $\gcd(a, b)$，**更核心的作用是同时求出裴蜀等式（Bézout's identity）$ax + by = \gcd(a, b)$ 中的整数解 $x$ 和 $y$**。

---

### 逻辑支撑 (CoT)

<details>
<summary>点击展开：ExGCD 的状态转移与数学推导</summary>

#### 1. 降维基石：标准欧几里得 (Standard GCD)
标准算法基于一个核心数学规律：两个数的最大公约数，等于其中较小的数与两数相除余数的最大公约数。
即：$\gcd(a, b) = \gcd(b, a \pmod b)$。
不断递归，直到 $b = 0$ 时，当前的 $a$ 就是最大公约数。

#### 2. 状态方程推导 (State Transition)
假设我们在进行除法运算 $a = q \cdot b + r$（其中 $q = \lfloor a/b \rfloor$，$r = a \pmod b$）。
我们要在每一步中维护等式：$ax + by = \gcd(a, b)$。

在进入下一层递归 $\gcd(b, a \pmod b)$ 时，假设我们已经求出了下一层的系数 $x_1$ 和 $y_1$，使得：
$b \cdot x_1 + (a \pmod b) \cdot y_1 = \gcd(a, b)$

将 $a \pmod b = a - q \cdot b$ 代入上式：
$b \cdot x_1 + (a - q \cdot b) \cdot y_1 = \gcd(a, b)$
$a \cdot y_1 + b \cdot (x_1 - q \cdot y_1) = \gcd(a, b)$

与原等式 $ax + by = \gcd(a, b)$ 对比系数，得出**绝对状态转移方程**：
*   $x = y_1$
*   $y = x1 - \lfloor a/b \rfloor \cdot y_1$

#### 3. 递归的终点 (The Base Case)
当 $b = 0$ 时，$\gcd(a, 0) = a$。
此时的裴蜀等式变为 $a \cdot x + 0 \cdot y = a$。
显然，解为 $x = 1, y = 0$。这是所有状态回溯的绝对起点。

</details>

---

### 裴蜀定理 - Bézout's identity
- 对于任意两个不全为零的整数 a 和 b，设它们的最大公约数为 `d = gcd(a,b)`，则必然存在整数 x 和 y，使得线性二元一次方程 `ax + by = d` 成立
- a 和 b 的任意线性组合 ax + by 的结果，必须是 gcd(a,b) 的整数倍

#### 在代码中的作用
- 乘法逆元是否存在
- 理论上保证了我们之前讨论的“扩展欧几里得算法 (ExGCD)”必定能求出解

#### 风险点
- 每次穷举循环前， 前置 target % gcd(A, B) == 0 的数学拦截
- 如果a，b允许负数，在传入 gcd(a, b) 函数前，必须执行绝对值映射 (abs(a), abs(b))，确保底层的辗转相除法在正整数域内安全收敛