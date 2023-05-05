import axios from "axios";

const setupClientSideAxiosInstance = () => {
  axios.defaults.baseURL = "http://localhost:8080";
};

const AxiosUtils = { setupClientSideAxiosInstance };
export default AxiosUtils;
