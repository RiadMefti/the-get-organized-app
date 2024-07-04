import { FC } from "react";
import { jwtDecode } from "jwt-decode";

interface DashboardProps {}

const Dashboard: FC<DashboardProps> = ({}) => {
  const token = localStorage.getItem("token");
  const decodedData = jwtDecode(token!);
  console.log(decodedData); // Your JWT data is now accessible here
  return <div>Dashboard</div>;
};

export default Dashboard;
