package main

import (
        "fmt"
        "math"
        "strconv"
        "time"

)

//------------------------------------------------------------------------------
func Calc_Time_1(
        s_year  int, 
        s_month time.Month, 
        s_day   int, 
        s_hour  int, 
        s_min   int, 
        s_sec   int,
        e_year  int, 
        e_month time.Month, 
        e_day   int, 
        e_hour  int, 
        e_min   int, 
        e_sec   int,
        ) {

        p := fmt.Println

        //then := time.Date(e_year, e_month, e_day, e_hour, e_min, 0, time.UTC)
        //then := time.Date(1950,    9,       5,      3,     30,   0,  time.UTC)
        //then := time.Date(2018, 12, 9, 0, 34, 58, 651387237, time.UTC)

        // then := time.Date(e_year, e_month, e_day, 0, 34, 58, 651387237, time.UTC)
        //then := time.Date(e_year, e_month, e_day, 0, 34, 58, 651387237, time.UTC)

        start := time.Date(
        s_year  ,
        s_month ,
        s_day   ,
        s_hour  ,
        s_min   ,
        s_sec   ,
        0, 
        time.UTC,
        )

        end := time.Date(
        e_year  ,
        e_month ,
        e_day   ,
        e_hour  ,
        e_min   ,
        e_sec   ,
        0, 
        time.UTC,
        )

        p("Calc_Time_1 start =",start)
        p("Calc_Time_1 end   =",end  )

        year, month, day, hour, min, sec := diff(start, end)

        dff := end.Sub(start)
        //dff := start.Sub(end)
        p("dff =",dff)

        p("dff.Hours())               =",dff.Hours())
        p("dff.Minutes())             =",dff.Minutes())

        p("dff.Seconds())             =",dff.Seconds())
        p("int64(dff.Seconds()))      =",int64(dff.Seconds()))
        p("math.Round(dff.Seconds())) =",math.Round(dff.Seconds()))

        p("dff.Nanoseconds())          =",dff.Nanoseconds())



        fmt.Printf("Delta = %d years, %d months, %d days, %d hours, %d mins and %d seconds old.",
            year, month, day, hour, min, sec)


         afterTenSeconds := start.Add(time.Second * 10)
        afterTenMinutes := start.Add(time.Minute * 10)
        afterTenHours := start.Add(time.Hour * 10)
        afterTenDays := start.Add(time.Hour * 24 * 10)

        fmt.Printf("start = %v\n", start)
        fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
        fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
        fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
        fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)



} // func Calc_Time_1

//------------------------------------------------------------------------------
// func Calc_Time_2(
//         s_year, s_month, s_day, s_hour, s_min, s_sec 
//         e_year, e_month, e_day, e_hour, e_min, e_sec int) 
//         (year, month, day, hour, min, sec int) {
// } // func Calc_Time_2

//------------------------------------------------------------------------------
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}

//------------------------------------------------------------------------------
func plural(count int, singular string) (result string) {
        if (count == 1) || (count == 0) {
         result = strconv.Itoa(count) + " " + singular + " "
        } else {
         result = strconv.Itoa(count) + " " + singular + "s "
        }
 return
}

func secondsToHuman(input int) (result string) {
        years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
        seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
        months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
        seconds = input % (60 * 60 * 24 * 7 * 30)
        weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
        seconds = input % (60 * 60 * 24 * 7)
        days := math.Floor(float64(seconds) / 60 / 60 / 24)
        seconds = input % (60 * 60 * 24)
        hours := math.Floor(float64(seconds) / 60 / 60)
        seconds = input % (60 * 60)
        minutes := math.Floor(float64(seconds) / 60)
        seconds = input % 60

        if years > 0 {
                 result = plural(int(years), "year") + plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else if months > 0 {
                 result = plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else if weeks > 0 {
                 result = plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else if days > 0 {
                 result = plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else if hours > 0 {
                 result = plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else if minutes > 0 {
                 result = plural(int(minutes), "minute") + plural(int(seconds), "second")
        } else {
                 result = plural(int(seconds), "second")
        }

 return
}

func main_currentTime() {       
        fmt.Println("\n######################################\n")
        currentTime := time.Now()   
        fmt.Println("\n######################################\n")
             
        fmt.Println("currentTime",currentTime)          
        fmt.Println("\n######################################\n")
        fmt.Println(currentTime.Format("2006-01-02 15:04:05"))

        fmt.Println("\n######################################\n")   
        timeStampString := currentTime.Format("2006-01-02 15:04:05")    
        layOut := "2006-01-02 15:04:05"    
        timeStamp, err := time.Parse(layOut, timeStampString)
        if err != nil {
        fmt.Println(err)          
        }   
        hr, min, sec := timeStamp.Clock()

        fmt.Println("Year   :", currentTime.Year())
        fmt.Println("Month  :", currentTime.Month())
        fmt.Println("Day    :", currentTime.Day())
        fmt.Println("Hour   :", hr)
        fmt.Println("Min    :", min)
        fmt.Println("Sec    :", sec)    

        fmt.Println("\n######################################\n")   
        year, month, day := time.Now().Date()
        fmt.Println("Year   :", year)
        fmt.Println("Month  :", month)
        fmt.Println("Day    :", day)

        fmt.Println("\n######################################\n")          
        t := time.Now()

        y := t.Year()
        mon := t.Month()
        d := t.Day()
        h := t.Hour()
        m := t.Minute()
        s := t.Second()
        n := t.Nanosecond()

        fmt.Println("Year   :",y)
        fmt.Println("Month   :",mon)
        fmt.Println("Day   :",d)
        fmt.Println("Hour   :",h)
        fmt.Println("Minute :",m)
        fmt.Println("Second :",s)
        fmt.Println("Nanosec:",n)
}

//------------------------------------------------------------------------------
func main_1() {

        fmt.Println("\n######################################\n")   
        fmt.Println("100 seconds : ", secondsToHuman(100))
        fmt.Println("1000 seconds : ", secondsToHuman(1000))
        fmt.Println("3600 seconds : ", secondsToHuman(3600))
        fmt.Println("9999 seconds : ", secondsToHuman(9999))
        fmt.Println("8888888888 seconds : ", secondsToHuman(8888888888))


        fmt.Println("\n######################################\n")   
        birthday := time.Date(1950, 9, 5, 3, 30, 0, 0, time.UTC)
        year, month, day, hour, min, sec := diff(birthday, time.Now())

        fmt.Printf("You are %d years, %d months, %d days, %d hours, %d mins and %d seconds old.",
            year, month, day, hour, min, sec)

        fmt.Println("\n######################################\n")   
        main_currentTime()

        fmt.Println("\n######################################\n")   
        t1 := time.Now()
        t2 := t1.Add(time.Second * 100)

        fmt.Println(t1)
        fmt.Println(t2)

        diff := t2.Sub(t1).Seconds()

        fmt.Println(diff)

        // m, _:= time.ParseDuration (string(diff))
        // //m, _:= time.ParseDuration (diff)
        // fmt.Println("ParseDuration .Seconds",m.Seconds())          

        fmt.Println("\n######################################\n")   

        fmt.Println("\n######################################\n")   
        t1 = time.Date(1984, time.November, 3, 13, 0, 0, 0, time.UTC)
        t2 = time.Date(1984, time.November, 3, 10, 23, 34, 0, time.UTC)

        hs := t1.Sub(t2).Hours()

        hs, mf := math.Modf(hs)
        ms := mf * 60

        ms, sf := math.Modf(ms)
        ss := sf * 60

        fmt.Println(hs, "hours", ms, "minutes", ss, "seconds")
} // func main



//------------------------------------------------------------------------------
func main_2() {



    p := fmt.Println

    // We'll start by getting the current time.
    now := time.Now()
    p("now =",now)

    // You can build a `time` struct by providing the
    // year, month, day, etc. Times are always associated
    // with a `Location`, i.e. time zone.
    then := time.Date(2018, 12, 9, 0, 34, 58, 651387237, time.UTC)
    p("then =",then)

    // You can extract the various components of the time
    // value as expected.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // The Monday-Sunday `Weekday` is also available.
    p("then.Weekday() =",then.Weekday())

    // These methods compare two times, testing if the
    // first occurs before, after, or at the same time
    // as the second, respectively.
    p("then.Before(now))then =",then.Before(now))
    p("then.After(now)) then =",then.After(now))
    p("then.Equal(now)) then =",then.Equal(now))

    // The `Sub` methods returns a `Duration` representing
    // the interval between two times.
    diff := now.Sub(then)
    p("diff =",diff)

    // We can compute the length of the duration in
    // various units.
    p("diff.Hours())        =",diff.Hours())
    p("diff.Minutes())      =",diff.Minutes())
    p("diff.Seconds())      =",diff.Seconds())
    //p("diff.Round.Seconds())      =",diff.Round(diff.Seconds()))

        p("int64(diff.Seconds()))      =",int64(diff.Seconds()))
        p("math.Round(diff.Seconds())) =",math.Round(diff.Seconds()))

        //dsfs.Slot_Price                               = fmt.Sprintf("%f", ree_Price[n])

    p("diff.Nanoseconds())  =",diff.Nanoseconds())

    // You can use `Add` to advance a time by a given
    // duration, or with a `-` to move backwards by a
    // duration.
    p("then.Add(diff))  =",then.Add(diff))
    p("then.Add(-diff)) =",then.Add(-diff))

    //    fmt.Println("\n######################################\n")   
    //    t1 := time.Now()
    //    t2 := t1.Add(time.Second * 100)
    //
    //    fmt.Println(t1)
    //    fmt.Println(t2)
    //
    //    diff = t2.Sub(t1).Seconds()
    //
    //    fmt.Println(diff)

} // func main_x



//------------------------------------------------------------------------------
func main() {

        //p := fmt.Println
        Calc_Time_1(2018, 12, 15, 0, 0, 0,
                    2018, 12, 18, 0, 0, 0, )
} // func main

