import axios from "axios"

async function getCurrency(baseCoin: string, targetCoin: string): Promise<any> {
    const response = await axios.get(`http://localhost:5001/currency/${baseCoin}?targets=${targetCoin}`)
    return response.data
}

export { getCurrency }