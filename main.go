package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Skill struct {
	name     string
	progress float64
}

func main() {
	skills := make(map[string][]string)
	progress := make(map[string]float64)
	inDegree := make(map[string]int)
// file reading and parsing
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			parent, child := parts[0], parts[1]
			skills[parent] = append(skills[parent], child)
			inDegree[child]++
			if _, exists := inDegree[parent]; !exists {
				inDegree[parent] = 0
			}
		} else if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			skill := parts[0]
			var prog float64
			fmt.Sscanf(parts[1], "%f", &prog)
			progress[skill] = prog
			if _, exists := inDegree[skill]; !exists {
				inDegree[skill] = 0
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Topilogical  Sorting with Progress Consideration in subjects
	var result []string
	pq := make(PriorityQueue, 0)

	for skill, deg := range inDegree {
		if deg == 0 {
			push(&pq, Skill{name: skill, progress: progress[skill]})
		}
	}

	for len(pq) > 0 {
		skill := pop(&pq)
		result = append(result, skill.name)
		for _, child := range skills[skill.name] {
			inDegree[child]--
			if inDegree[child] == 0 {
				push(&pq, Skill{name: child, progress: progress[child]})
			}
		}
	}

	//printing the result
	for _, skill := range result {
		fmt.Println(skill)
	}
}

type PriorityQueue []Skill

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].progress > pq[j].progress || (pq[i].progress == pq[j].progress && pq[i].name < pq[j].name)
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func push(pq *PriorityQueue, item Skill) {
	*pq = append(*pq, item)
	sort.Sort(pq)
}

func pop(pq *PriorityQueue) Skill {
	item := (*pq)[0]
	*pq = (*pq)[1:]
	return item
}
