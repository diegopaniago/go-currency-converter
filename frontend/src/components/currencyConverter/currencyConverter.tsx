import React, {  useState } from 'react';
import { getCurrency } from '../../apis/currencyApi';
import './currencyConverter.css';

async function convert(baseCoin: string, targetCoin: string, setCurrency: any): Promise<void> {
  try {
    const response = await getCurrency(baseCoin, targetCoin);
    setCurrency(JSON.stringify(response));
  } catch (error) {
    alert('Error while converting currency!');
    console.error(error);
  }
}

function typeHandler(event: React.ChangeEvent<HTMLInputElement>, callback: any) {
  event.target.value = event.target.value.toUpperCase();
  callback(event.target.value);
}


const CurrencyConverter: React.FC = () => {
  const [baseCoin, setBaseCoin] = useState('');
  const [targetCoin, setTargetCoin] = useState('');
  const [currency, setCurrency] = useState('');
  
  return (
      <div className='currency-box'>
        <div>
          <label>Base coin:</label>
          <input type='text' onChange={(event) => typeHandler(event, setBaseCoin)}/>
        </div>
        <div>
          <label>Target coin:</label>
          <input type='text' onChange={(event) => typeHandler(event, setTargetCoin)} />
        </div>
        <div>
          <button onClick={() => convert(baseCoin, targetCoin, setCurrency)}>Convert</button>
        </div>
        <textarea disabled value={currency}></textarea>
      </div>
    );
  }

export default CurrencyConverter;