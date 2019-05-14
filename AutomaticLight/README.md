Problem A
Automatic Light
A company has an automatic light at the front door. When an employee enters the building the light lights up immediately and stays on for the next x minutes. If the light is already on, it remains on and reprograms itself to stay on for the next x minutes. Notice that when an employee leaves the building, the light isn’t affected.

You have the list of times when each employee will enter or leave the building and want to configure the light with a new value x in order to save energy.

Find the smallest number x that will leave no employees in the dark, that is, while there is someone at the company, the lights won’t go out.

Input
The first line of input will consist of one integer n (1≤n≤103), the number of employees.

n lines will follow. The i-th line will contain two integers, tin and tout (0≤tin<tout≤40000), the time, in minutes, that the i-th employee enters and leaves the building, respectively.

Each employee will already be in the building at the start of minute tin and will already be outside at the start of minute tout, that is, the light needs to be on at minute tin, but not necessarily at minute tout.

Output
Print a single integer, the smallest value x that will leave no employees in the dark.

Sample Input 1	Sample Output 1
1
5 278
273
Sample Input 2	Sample Output 2
4
0 10
4 6
11 14
12 15
6

https://vtexcc19.kattis.com/problems/vtexcc19.automaticlight