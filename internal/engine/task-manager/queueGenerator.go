package task_manager

import (
	"github.com/thisismeamir/kage/internal/internal-pkg/flow"
	"log"
	"sort"
)

func (tm *TaskManager) QueueGenerator() {
	log.Printf("\n\n\n Queue Generator Started \n\n\n ")
	nodes, maps := tm.GetAllTasks().GroupByLevel().Sort().SplitGroupedTaskByType()
	log.Printf("Queuing %v ", nodes)
	log.Printf("Queuing %v", maps)

}

type GroupedTask map[int][]flow.TaskRegister

type ListOfTaskRegisters []flow.TaskRegister

func (tm TaskManager) GetAllTasks() ListOfTaskRegisters {
	var allTasks []flow.TaskRegister
	for _, fl := range tm.Flows {
		currentFlow := flow.LoadFlow(fl.Path)
		allTasks = append(allTasks, currentFlow.Tasks...)
	}
	log.Printf("\n\n\nQueuing %v\n\n\n ", allTasks)
	return allTasks
}

func (tasks ListOfTaskRegisters) GroupByLevel() GroupedTask {
	groupedByLevel := make(map[int][]flow.TaskRegister)
	for _, task := range tasks {
		groupedByLevel[task.Level] = append(groupedByLevel[task.Level], task)
	}
	log.Printf("Grouped Tasks Successfully %v", groupedByLevel)
	return groupedByLevel
}

func (groupedTasks GroupedTask) Sort() GroupedTask {
	for level, tasks := range groupedTasks {
		sort.SliceStable(tasks, func(i, j int) bool {
			return tasks[i].Queue < tasks[j].Queue
		})

		sort.SliceStable(tasks, func(i, j int) bool {
			return tasks[i].Urgency > tasks[j].Urgency
		})

		groupedTasks[level] = tasks
	}
	log.Printf("Sorted GroupedTasks %v", groupedTasks)
	return groupedTasks
}

func (groupedTask GroupedTask) SplitGroupedTaskByType() (GroupedTask, GroupedTask) {
	nodes := make(GroupedTask)
	maps := make(GroupedTask)

	for level, tasks := range groupedTask {
		var nodeTasks []flow.TaskRegister
		var mapTasks []flow.TaskRegister

		for _, task := range tasks {
			switch task.Type {
			case "node":
				nodeTasks = append(nodeTasks, task)
			case "map":
				mapTasks = append(mapTasks, task)
			default:
				continue
			}
		}
		if len(nodeTasks) > 0 {
			nodes[level] = nodeTasks
		}
		if len(mapTasks) > 0 {
			maps[level] = mapTasks
		}
	}
	log.Printf("Node Tasks %v", nodes)
	log.Printf("Map Tasks %v", maps)
	return nodes, maps
}
