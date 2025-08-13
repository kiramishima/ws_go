package main

import "fmt"

func GroupOpenHours(openHours []map[string]string) []map[string]string {
	var daysmap = make(map[string]map[string]string)
	var grouped = make(map[string][]map[string]string)
	days := map[string]string{
		"Monday":    "Tuesday",
		"Tuesday":   "Wednesday",
		"Wednesday": "Thursday",
		"Thursday":  "Friday",
		"Friday":    "Saturday",
		"Saturday":  "Sunday",
		"Sunday":    ""}

	for _, v := range openHours {
		v["next"] = days[v["day"]]
		daysmap[v["day"]] = v
	}

	for v := range days {
		if _, ok := daysmap[v]; !ok {
			daysmap[v] = map[string]string{"open": "", "close": "", "day": v}
		}
	}
	fmt.Println(daysmap)

	for _, v := range daysmap {

		key := fmt.Sprintf("%s-%s", v["open"], v["close"])
		grouped[key] = append(grouped[key], v)
	}
	fmt.Print(grouped)

	// fmt.Println(keys)
	for k, v := range grouped {
		fmt.Println(k, len(v))
		for idx, item := range v {
			fmt.Println(idx, item)
			if v[idx]["next"] == v[idx+1]["day"] {

			}
		}
	}
	/*var m = make(map[string][]string)
	for k := range keys {
		// var pivot =
		var val = days[k]
		var current = daysmap[val]
		var cnt = k
		// fmt.Println(current["day"])
	out:
		for i := 0; i < len(dayz); i++ {
			day := dayz[i]
			nextDay := daysmap[day]
			fmt.Println(current["day"], nextDay["day"])
			if current["day"] == nextDay["day"] {
				continue
			}
			// fmt.Println(current["day"], nextDay["day"], "cnt", cnt, "i", i)
			if current["open"] == nextDay["open"] && current["close"] == nextDay["close"] {
				if i-cnt == 1 {
					key := fmt.Sprintf("%s-%s", current["day"], nextDay["day"])
					m[key] = append(m[key], current["day"], nextDay["day"])
				} else {
					break out
				}
				cnt++
			} else {
				key := nextDay["day"]
				m[key] = append(m[key], nextDay["day"])
			}
		}
	}
	fmt.Println(m)
	//
	/*var grouped = make(map[string][]string)
	for k, v := range days {
		// fmt.Println(daysmap[k], daysmap[v])
		if k == "" || v == "" {
			// key := fmt.Sprintf("%s-%s", daysmap[k]["open"], daysmap[k]["close"])
			grouped[k] = append(grouped[k], k)
		} else if v != "" && (daysmap[k]["open"] != daysmap[v]["open"] || daysmap[k]["close"] != daysmap[v]["close"]) {
			fmt.Println(v)

			grouped[k] = append(grouped[k], k)
		} else if (v != "") && (daysmap[k]["open"] == daysmap[v]["open"] && daysmap[k]["close"] == daysmap[v]["close"]) {
			key := fmt.Sprintf("%s-%s", k, v)
			delete(daysmap, k)
			grouped[key] = append(grouped[key], k, v)
		}
	}
	fmt.Println(grouped)*/
	return make([]map[string]string, 0)
}

func main() {
	openHours := []map[string]string{
		{
			"day":   "Monday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Tuesday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Wednesday",
			"open":  "8:00 AM",
			"close": "6:00 PM",
		},
		{
			"day":   "Thursday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Friday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Saturday",
			"open":  "8:00 AM",
			"close": "4:00 PM",
		},
	}

	GroupOpenHours(openHours)
}
