import axios from "axios";

const API = process.env.REACT_APP_API_URL;

// queryString contains symbols (splitted by comma) and range
export async function requestChart(queryString) {
  try {
    const resp = await axios.get(`${API}/stocks${queryString}`);
    return resp.data;
  } catch (err) {
    if (err.response) {
      throw err.response.data;
    } else if (err.request) {
      throw err.request;
    } else {
      throw err.message;
    }
  }
}
