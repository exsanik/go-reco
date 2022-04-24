import axios from "axios";

import { BASE_URL } from "src/config";

export const apiInstance = axios.create({ baseURL: BASE_URL });
