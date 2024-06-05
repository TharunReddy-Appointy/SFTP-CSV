package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// ReportDataV4 represents the structure of your report data
type ReportDataV4 struct {
	BookingStatus          string
	BookingDate            time.Time
	StartTime              time.Time
	DeliveryMedium         string
	Category               string
	Duration               int
	ServiceType            string
	ServiceName            string
	BookingType            string
	ChannelName            string
	StaffFirstName         string
	StaffLastName          string
	StaffEmail             string
	StaffRole              string
	CustomerEmail          string
	BuildingCode           string
	RoomName               string
	RoomCode               string
	BookingId              string
	ClassId                string
	BookedByEmail          string
	UpdatedDateTime        time.Time
	UpdatedBy              string
	CostCategory           string
	CostTier               string
	AmountPaid             float64
	AmountRefunded         float64
	ReminderSentYN         string
	ItemType               string
	CurrencyType           string
	PaymentStatus          string
	PaymentTimestamp       time.Time
	PaymentRefundDate      time.Time
	PaymentReason          string
	Timezone               string
	CapacityConfirmed      int64
	CapacityWaitlist       int64
	UserDepartment         string
	UserBusiness           string
	DivisionConpanyCode    string
	LocalStartTime         string
	DayOfWeek              string
	LocationCode           string
	ShiftId                string
	ExternalAttendeesCount int64
}

func main() {
	start := time.Now()

	// Generate sample CSV data
	err := generateAndWriteCSV(1000000)
	if err != nil {
		log.Fatalf("failed to generate and write CSV: %v", err)
	}
	generateAndWriteTime := time.Since(start)
	fmt.Printf("Time taken to generate and write CSV: %v\n", generateAndWriteTime)

	// Get the file size
	fileInfo, err := os.Stat("sample.csv")
	if err != nil {
		log.Fatalf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()
	fmt.Printf("Size of the generated CSV file: %d bytes\n", fileSize)

	// Read the CSV data back into memory
	start = time.Now()
	data, err := readCSVData("sample.csv")
	if err != nil {
		log.Fatalf("failed to read CSV data: %v", err)
	}
	readTime := time.Since(start)
	fmt.Printf("Time taken to read CSV data: %v\n", readTime)

	fmt.Printf("CSV data successfully generated and read. Total records: %d\n", len(data))

	duration := time.Since(start)
	fmt.Printf("Total time taken: %v\n", duration)
}

func generateAndWriteCSV(numEntries int) error {
	// Generate sample CSV data
	err := generateSampleCSV(numEntries)
	if err != nil {
		return err
	}

	// Write the generated CSV file to the local system
	err = writeCSVToLocal("sample.csv")
	if err != nil {
		return err
	}

	return nil
}

func generateSampleCSV(numEntries int) error {
	file, err := os.Create("sample.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{
		"BookingStatus",
		"BookingDate",
		"StartTime",
		"DeliveryMedium",
		"Category",
		"Duration",
		"ServiceType",
		"ServiceName",
		"BookingType",
		"ChannelName",
		"StaffFirstName",
		"StaffLastName",
		"StaffEmail",
		"StaffRole",
		"CustomerEmail",
		"BuildingCode",
		"RoomName",
		"RoomCode",
		"BookingId",
		"ClassId",
		"BookedByEmail",
		"UpdatedDateTime",
		"UpdatedBy",
		"CostCategory",
		"CostTier",
		"AmountPaid",
		"AmountRefunded",
		"ReminderSentYN",
		"ItemType",
		"CurrencyType",
		"PaymentStatus",
		"PaymentTimestamp",
		"PaymentRefundDate",
		"PaymentReason",
		"Timezone",
		"CapacityConfirmed",
		"CapacityWaitlist",
		"UserDepartment",
		"UserBusiness",
		"DivisionConpanyCode",
		"LocalStartTime",
		"DayOfWeek",
		"LocationCode",
		"ShiftId",
		"ExternalAttendeesCount",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data
	for i := 0; i < numEntries; i++ {
		record := []string{
			"confirmed",
			time.Now().Format(time.RFC3339),
			time.Now().Format(time.RFC3339),
			"Medium" + strconv.Itoa(i),
			"Category" + strconv.Itoa(i),
			strconv.Itoa(i),
			"Type" + strconv.Itoa(i),
			"Name" + strconv.Itoa(i),
			"Type" + strconv.Itoa(i),
			"Channel" + strconv.Itoa(i),
			"First" + strconv.Itoa(i),
			"Last" + strconv.Itoa(i),
			"email" + strconv.Itoa(i) + "@example.com",
			"Role" + strconv.Itoa(i),
			"customer" + strconv.Itoa(i) + "@example.com",
			"Code" + strconv.Itoa(i),
			"Room" + strconv.Itoa(i),
			"Code" + strconv.Itoa(i),
			"ID" + strconv.Itoa(i),
			"Class" + strconv.Itoa(i),
			"booked" + strconv.Itoa(i) + "@example.com",
			time.Now().Format(time.RFC3339),
			"UpdatedBy" + strconv.Itoa(i),
			"Category" + strconv.Itoa(i),
			"Tier" + strconv.Itoa(i),
			strconv.FormatFloat(float64(i)*0.5, 'f', -1, 64),
			strconv.FormatFloat(float64(i)*0.3, 'f', -1, 64),
			"Y",
			"Item" + strconv.Itoa(i),
			"Currency",
			"Status" + strconv.Itoa(i),
			time.Now().Format(time.RFC3339),
			time.Now().Format(time.RFC3339),
			"Reason" + strconv.Itoa(i),
			"Timezone" + strconv.Itoa(i),
			strconv.Itoa(i),
			strconv.Itoa(i),
			"Department" + strconv.Itoa(i),
			"Business" + strconv.Itoa(i),
			"DivisionCode" + strconv.Itoa(i),
			"StartTime" + strconv.Itoa(i),
			"Day" + strconv.Itoa(i),
			"Code" + strconv.Itoa(i),
			"Shift" + strconv.Itoa(i),
			strconv.Itoa(i),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func writeCSVToLocal(filename string) error {
	return nil
}

func readCSVData(filename string) ([]ReportDataV4, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var data []ReportDataV4

	// Read header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read data
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Parse data
		d := ReportDataV4{
			BookingStatus:          record[0],
			BookingDate:            parseTime(record[1]),
			StartTime:              parseTime(record[2]),
			DeliveryMedium:         record[3],
			Category:               record[4],
			Duration:               parseInt(record[5]),
			ServiceType:            record[6],
			ServiceName:            record[7],
			BookingType:            record[8],
			ChannelName:            record[9],
			StaffFirstName:         record[10],
			StaffLastName:          record[11],
			StaffEmail:             record[12],
			StaffRole:              record[13],
			CustomerEmail:          record[14],
			BuildingCode:           record[15],
			RoomName:               record[16],
			RoomCode:               record[17],
			BookingId:              record[18],
			ClassId:                record[19],
			BookedByEmail:          record[20],
			UpdatedDateTime:        parseTime(record[21]),
			UpdatedBy:              record[22],
			CostCategory:           record[23],
			CostTier:               record[24],
			AmountPaid:             parseFloat(record[25]),
			AmountRefunded:         parseFloat(record[26]),
			ReminderSentYN:         record[27],
			ItemType:               record[28],
			CurrencyType:           record[29],
			PaymentStatus:          record[30],
			PaymentTimestamp:       parseTime(record[31]),
			PaymentRefundDate:      parseTime(record[32]),
			PaymentReason:          record[33],
			Timezone:               record[34],
			CapacityConfirmed:      parseInt64(record[35]),
			CapacityWaitlist:       parseInt64(record[36]),
			UserDepartment:         record[37],
			UserBusiness:           record[38],
			DivisionConpanyCode:    record[39],
			LocalStartTime:         record[40],
			DayOfWeek:              record[41],
			LocationCode:           record[42],
			ShiftId:                record[43],
			ExternalAttendeesCount: parseInt64(record[44]),
		}

		data = append(data, d)
	}

	return data, nil
}

func parseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}
	return t
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v", err)
	}
	return i
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("failed to parse float: %v", err)
	}
	return f
}

func parseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("failed to parse int64: %v", err)
	}
	return i
}
