package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupOpenHours", func() {
	It("should work on the provided example", func() {
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
		expected := []map[string]string{
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
		actual := GroupOpenHours(openHours)
		Expect(actual).To(Equal(expected))
	})

	It("should handle ranges of closed days", func() {
		openHours := []map[string]string{
			{
				"day":   "Monday",
				"open":  "9:00 AM",
				"close": "4:00 PM",
			},
			{
				"day":   "Tuesday",
				"open":  "9:00 AM",
				"close": "4:00 PM",
			},
			{
				"day":   "Wednesday",
				"open":  "9:00 AM",
				"close": "4:00 PM",
			},
			{
				"day":   "Thursday",
				"open":  "9:00 AM",
				"close": "4:00 PM",
			},
		}
		expected := []map[string]string{
			{
				"days":  "Monday-Thursday",
				"open":  "9:00 AM",
				"close": "4:00 PM",
			},
			{
				"days":  "Friday-Sunday",
				"open":  "",
				"close": "",
			},
		}
		actual := GroupOpenHours(openHours)
		Expect(actual).To(Equal(expected))
	})
})
