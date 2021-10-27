import { RoomsInterface } from "./IRoom";
import { CustomersInterface } from "./ICustomer";
import { RoomPaymentsInterface } from "./IRoomPayment";
import { EmployeesInterface } from "./IEmployee";


export interface CheckInInterface {
  ID: number,
  DateTime: Date,  

	ReserveID: number,
	Reserve: RoomsInterface,
	PaymentID: number,
	Payment:   RoomPaymentsInterface,
	EmployeeID: number,
	Employee: EmployeesInterface,
    CheckInID: number,
	CheckIn:   CustomersInterface,
}