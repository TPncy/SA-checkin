import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { EmployeesInterface } from "../models/IEmployee";
import { CustomersInterface } from "../models/ICustomer";
import { RoomsInterface } from "../models/IRoom";
import { RoomPaymentsInterface } from "../models/IRoomPayment";
import { CheckInInterface } from "../models/ICheckIn";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function CheckInCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [employees, setEmployees] = useState<EmployeesInterface>();
  const [customers, setCustomers] = useState<CustomersInterface[]>([]);
  const [rooms, setRooms] = useState<RoomsInterface[]>([]);
  const [roompayments, setRoomPayments] = useState<RoomPaymentsInterface[]>([]);
  const [checkIn, setCheckIn] = useState<Partial<CheckInInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}`,
    "Content-Type": "application/json",
  },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof checkIn;
    setCheckIn({
      ...checkIn,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getEmployee = async () => {
    const uid = Number(localStorage.getItem("uid"));
    fetch(`${apiUrl}/employee/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setEmployees(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getCustomer = async () => {
    fetch(`${apiUrl}/customers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCustomers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRoom = async () => {
    fetch(`${apiUrl}/rooms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRooms(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRoomPayment = async () => {
    fetch(`${apiUrl}/roompayments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRoomPayments(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getEmployee();
    getCustomer();
    getRoom();
    getRoomPayment();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      CheckInID: convertType(checkIn.CheckInID),
      ReserveID: convertType(checkIn.ReserveID),
      PaymentID: convertType(checkIn.PaymentID),
      EmployeeID:convertType(employees?.ID),
      DateTime: selectedDate,
    };

    const requestOptionsPost = {
      method: "POST",
      headers: { Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/check_ins`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
        }
      });
  }
    return (

        <Container className={classes.container} maxWidth="md">
        <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="success">
            บันทึกข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            บันทึกข้อมูลไม่สำเร็จ
          </Alert>
        </Snackbar>
        <Paper className={classes.paper}>
          <Box display="flex">
            <Box flexGrow={1}>
              <Typography
                component="h2"
                variant="h6"
                color="primary"
                gutterBottom
              >
                บันทึกการเข้าห้องพัก
              </Typography>
            </Box>
          </Box>
          <Divider />
          <Grid container spacing={3} className={classes.root}>
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>หมายเลขการชำระเงิน</p>
                <Select
                  native
                  value={checkIn.PaymentID}
                  onChange={handleChange}
                  inputProps={{
                    name: "PaymentID",
                  }}
                > 
                  <option aria-label="None" value="">
                    กรุณาเลือกหมายเลขการชำระเงิน
                  </option>
                  {roompayments.map((item: RoomPaymentsInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.ID}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>หมายเลขห้อง</p>
                <Select
                  native
                  value={checkIn.ReserveID}
                  onChange={handleChange}
                  inputProps={{
                    name: "ReserveID",
                  }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกหมายเลขห้อง
                  </option>
                  {rooms.map((item: RoomsInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Roomnumber}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>รหัสลูกค้า</p>
                <Select
                  native
                  value={checkIn.CheckInID}
                  onChange={handleChange}
                  inputProps={{
                    name: "CheckInID",
                  }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกรหัสลูกค้า
                  </option>
                  {customers.map((item: CustomersInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.ID}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>พนักงานที่บันทึกข้อมูล</p>
                <Select
                  native
                  disabled
                  value={checkIn.EmployeeID}
                  /*onChange={handleChange}
                  inputProps={{
                    name: "EmployeeID",
                  }}*/
                >
                  <option aria-label="None" value="">
                  {employees?.Name}
                </option>
                
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>วัน/เวลาในการดำเนินการ</p>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                  <KeyboardDateTimePicker
                    name="Date_time"
                    value={selectedDate}
                    onChange={handleDateChange}
                    label="กรุณาเลือกวันที่และเวลา"
                    minDate={new Date("2018-01-01T00:00")}
                    format="yyyy/MM/dd hh:mm a"
                  />
                </MuiPickersUtilsProvider>
              </FormControl>
            </Grid>
            <Grid item xs={12}>
              <Button
                style={{ float: "right" }}
                variant="contained"
                onClick={submit}
                color="primary"
              >
                บันทึก
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    );
  }
  
  export default CheckInCreate;