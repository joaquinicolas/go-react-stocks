import axios from "axios";

const API = process.env.REACT_APP_API_URL;

// Range could be either 1 year or 1 month
export async function bySymbol(symbol, range = '1m') {
  try {
    const resp = await axios.get(`${API}/${symbol}/chart/${range}`);
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
