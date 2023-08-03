import { useContext, useEffect, useState } from "react";

import { Container, Grid } from "@mui/material";

import { MonthElement } from "@/components/templates/MonthElement";
import { createCalender } from "@/libs/service/calender";
import { MonthContext } from "@/provider/CalendarProvider";
import { margeSchedules } from "@/libs/service/schedule";

export const MonthCalender = () => {
  const { month, schedules, setDaySelected, setShowAddDialog } =
    useContext(MonthContext);
  const [currentMonth, setCurrentMonth] = useState(createCalender());
  const [calendar, setCalendar] = useState(
    margeSchedules(currentMonth, schedules)
  );

  useEffect(() => {
    const newCalendar = createCalender(month);
    setCurrentMonth(newCalendar);
    setCalendar(margeSchedules(newCalendar, schedules));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [month, schedules]);

  const days = ["日", "月", "火", "水", "木", "金", "土"];

  return (
    <div>
      <Container sx={{ marginTop: "10px" }}>
        <Grid container columns={{ xs: 7, sm: 7, md: 7 }}>
          {days.map((day) => (
            <Grid
              item
              xs={1}
              sm={1}
              md={1}
              key={day}
              sx={{
                borderBottom: "1px solid #ccc",
                textAlign: "center",
                fontWeight: "bold",
              }}
            >
              {day}
            </Grid>
          ))}
        </Grid>
        <Grid
          container
          columns={{ xs: 7, sm: 7, md: 7 }}
          sx={{ borderLeft: "1px solid #ccc" }}
        >
          {calendar.map((item: any, index: number) => (
            <Grid
              item
              xs={1}
              sm={1}
              md={1}
              key={index}
              sx={{
                borderRight: "1px solid #ccc",
                borderBottom: "1px solid #ccc;",
                textAlign: "right",
                width: "40px",
                height: "100px",
              }}
            >
              <div
                onClick={(e) => {
                  e.stopPropagation();
                  setDaySelected(item.date);
                  setShowAddDialog(true);
                }}
              >
                <MonthElement
                  key={index}
                  day={item.date}
                  schedule={item.schedules}
                />
              </div>
            </Grid>
          ))}
        </Grid>
      </Container>
    </div>
  );
};
