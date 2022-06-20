package utils

import (
	"sort"
	"strconv"
)

func ListToTree(list []map[string]interface{}, idKey, pidKey, childrenKey string, isGenId bool) (tree []map[string]interface{}) {
	listLen := len(list)
	if listLen == 0 {
		return
	}

	var idMap = make(map[string]map[string]interface{}, listLen)
	var pidMap = make(map[string]string, listLen)

	var ids []string
	var sortedList []map[string]interface{}

	for _, m := range list {
		id := (m[idKey]).(string)
		ids = append(ids, id)

		idMap[id] = m
		pid := (m[pidKey]).(string)
		pidMap[pid] = pid
	}

	for _, m := range list {
		id := (m[idKey]).(string)
		if pidMap[id] != "" {
			delete(pidMap, id)
		}
	}

	sort.Slice(ids, func(i, j int) bool {
		numA, _ := strconv.Atoi(ids[i])
		numB, _ := strconv.Atoi(ids[j])
		return numA < numB
	})

	for _, id := range ids {
		sortedList = append(sortedList, idMap[id])
	}

	for _, m := range sortedList {
		pid := (m[pidKey]).(string)
		if isGenId {
			m[idKey] = SnowflakeId(RandomInt(1024)).String()
		}
		if pidMap[pid] != "" {
			tree = append(tree, m)
		} else {
			if isGenId {
				id := (m[idKey]).(string)
				if pidMap[id] != "" {
					idMap[id][pidKey] = m[idKey]
				}
			}
			parent := idMap[pid]
			if parent[childrenKey] == nil {
				parent[childrenKey] = []map[string]interface{}{}

			}
			if pid != "0" {
				m[pidKey] = parent[idKey]
			}
			children := parent[childrenKey].([]map[string]interface{})
			children = append(children, m)
			parent[childrenKey] = children
		}
	}

	return
}
