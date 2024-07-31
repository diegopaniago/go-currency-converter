import React, {  useState } from "react";
import { getCurrency } from "../../apis/currencyApi";

async function convert(baseCoin: string, targetCoin: string, setCurrency: any): Promise<void> {
  const response = await getCurrency(baseCoin, targetCoin);
  setCurrency(JSON.stringify(response));
}


const CurrencyConverter: React.FC = () => {
  const [baseCoin, setBaseCoin] = useState("");
  const [targetCoin, setTargetCoin] = useState("");
  const [currency, setCurrency] = useState("");
  
  return (
      <div>
        <label>Base coin:</label>
        <input type="text" onChange={(event) => setBaseCoin(event?.target?.value || '')}/>
        <label>Target coin:</label>
        <input type="text" onChange={(event) => setTargetCoin(event?.target?.value || '')} />
        <button onClick={() => convert(baseCoin, targetCoin, setCurrency)}>Convert</button>
        <textarea value={currency}></textarea>
      </div>
    );
  }

export default CurrencyConverter;