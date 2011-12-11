#include <stdio.h>

#define MAX 100

int cheese[MAX];

int main()
{
  int n, i, average, total = 0, done = 0, lack = 0, move = 0;
/* 整数done表示前面有done个数据可以调整好了，lack表示缺失的奶酪数，move表示移动次数 */

  scanf("%d", &n);
  for (i = 0; i < n; i++) {
    scanf("%d", &cheese[i]);
    total += cheese[i];
  }
  average = total/n;
  for (i = 0; i < n-1; i++) {   /*  */
    if (done == i) {  /* 前面有i堆奶酪都可以调整好了 */
      if (cheese[i] < average)  
	lack += average - cheese[i]; /* cheese[i]需要从后面的堆转移奶酪给它 */ 
      else {   /* 前面有i堆奶酪分配好了，且cheese[i]有多余的酪，需转移到cheese[i+1] */
	done++;
	move++;
	cheese[i+1] += cheese[i] - 10;
      }
    }
    else if (cheese[i] >= lack) { 
/*要总数最少，必须从cheese[i]中一次性转移所需的奶酪数到cheese[i-1]，（否则暂不转移）,再把cheese[i-1]中多余的转移到cheese[i-2](用[i-1]->[i-2]表示)，然后[i-2]->[i-3]....[done+1]->[done]，总共转移i-done次,注：这种转移事实上是没有发生，只知道可以这么调整罢了 */
      cheese[i] -= lack;          
      move += i - done; 
      done = i--;   
    }
    else
      lack += average - cheese[i];/* 暂不转移 */
  }
  if (cheese[i] > average)  /* n特殊处理 */
    move += i - done;

  printf("%d\n", move);

}
