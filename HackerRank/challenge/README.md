# Task

Your team is working on a project for a retail chain which keeps their stores' weekly open hours listings as an array of data. The data is in the following format:

```
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
```

However, the company's website needs the data to be transformed to a grouped format for displaying to visitors. The grouped format is as follows:

```
groupedOpenHours := []map[string]string{
  {
    "days":  "Monday-Tuesday",
    "open":  "8:00 AM",
    "close": "5:00 PM",
  },
  {
    "days":  "Wednesday",
    "open":  "8:00 AM",
    "close": "6:00 PM",
  },
  {
    "days":  "Thursday-Friday",
    "open":  "8:00 AM",
    "close": "5:00 PM",
  },
  {
    "days":  "Saturday",
    "open":  "8:00 AM",
    "close": "4:00 PM",
  },
  {
    "days":  "Sunday",
    "open":  "",
    "close": "",
  },
}
```

In the output above, any consecutive days sharing the same open and close times have been compressed into the same map. Whenever consecutive days sharing open and close times are encountered, the first and last day in the range are concatenated with a hyphen for the `"day"` key.

Your task is to write code to perform the transformation. The function you'll complete is `GroupOpenHours(openHours []map[string]string) []map[string]string`. The function should return the transformed map slice in the above format.

The output slice should be in order of the days of the week. Consider Monday as the beginning of the week and Sunday as the end. No range that bridges the gap between Sunday-Monday should be created (but a range from Monday-Sunday is valid whenever the entire week has the same open/closed hours or the input is empty.

As shown above, any missing days of the week should be added to the returned slice as maps with `"open": ""` and `"close": ""` entries. When the missing dates consist of consecutive ranges, they should use the same hyphenated grouped `"days"` key format as open days would be.

The `openHours` parameter will always be well-formed but may be empty and unsorted. There will never be more than 7 maps in the slice, and each map is guaranteed to have only `"day"`, `"open"` and `"close"` keys present with string values formatted as in the structure shown above. All `"day"` keys are guaranteed to be unique in `openHours` and correctly capitalized, valid days of the week.

Please do not mutate the `openHours` parameter.