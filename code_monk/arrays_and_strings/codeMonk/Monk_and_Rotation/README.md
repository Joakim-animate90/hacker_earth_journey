## Monk and Rotation
Monk loves to preform different operations on arrays, and so being the principal of Hackerearth School, he assigned a task to his new student Mishki. Mishki will be provided with an integer array A of size N and an integer K , where she needs to rotate the array in the right direction by K steps and then print the resultant array. As she is new to the school, please help her to complete the task.

Video approach to solve this question: https://youtube.com/watch?v=mX38pWM--0M&list=PL1YS4hYJip07A-YteNUR8qTeA_wHQarDX&index=42

### Input:
The first line will consists of one integer T denoting the number of test cases.
For each test case:
1) The first line consists of two integers N and K, N being the number of elements in the array and K denotes the number of steps of rotation.
2) The next line consists of N space separated integers , denoting the elements of the array A.

### Output:
Print the required array.

### Constraints
1 <= T <= 20
1 <= N <= 10^5
0 <= K <= 10^6
0 <= A[i] <= 10^6


### Sample Input
1
5 2
1 2 3 4 5

### Sample Output
4 5 1 2 3

### Explanation
Here T is 1, which means one test case.
denoting the number of elements in the array and , denoting the number of steps of rotations.
The initial array is:
In first rotation, 5 will come in the first position and all other elements will move to one position ahead from their current position. Now, the resultant array will be
In second rotation, 4 will come in the first position and all other elements will move to one position ahead from their current position. Now, the resultant array will be


Time Limit: 1.0 sec(s) for each input file
Memory Limit: 256 MB
Source Limit: 1024 KB


### Result
```
Input			 Result	  Time (sec)	Memory (KiB)	Score	Your Output	Correct Output
Input #1	 Accepted	0.009602		2							10		
Input #2	 Accepted	0.009565		2							10		
Input #3	 Accepted	0.009119		2							20		
Input #4	 Accepted	0.017087		102984				20		
Input #5	 Accepted	0.041248		104712				20		
Input #6	 Accepted	0.009205		2							20	
```