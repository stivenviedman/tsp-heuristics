## TSP CLI

I've been always fascinated by the conceptual simplicity of the Traveling Salesman Problem and by the complexity of its solutions. In this small side project I try to implement some well-known algorithms for the problem as I use it to learn Go.

## Implementation of TSP algorithms

- Nearest Neighbor
- Farthest Insertion
- Greedy Algorithm (TODO)
- Nearest Insertion (TODO)
- Double Minimum Spanning Tree (TODO)

## How to use it

Modify the <code>tsp-data.json</code> file depending on your needs. The file expects two inputs: the name of the algorithm you want to use and the edges of the problem instance.

Execute <code>go run main/main.go</code> and the result should look like:

<code>[{1 1} {0 0} {3 3} {20 20}]</code>

<code>24.041630560342615</code>

Indicating the order of the route and the total distance.