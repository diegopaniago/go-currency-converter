import React, {  useState } from 'react';
import { getCurrency } from '../../apis/currencyApi';

async function convert(baseCoin: string, targetCoin: string, setCurrency: any): Promise<void> {
  try {
    const response = await getCurrency(baseCoin, targetCoin);
    console.log('res:', response)
    setCurrency(response);
  } catch (error) {
    alert('Error while converting currency!');
    console.error(error);
  }
}

function typingHandler(event: React.ChangeEvent<HTMLInputElement>, callback: any) {
  event.target.value = event.target.value.toUpperCase();
  callback(event.target.value);
}


const CurrencyConverter: React.FC = () => {
  const [baseCoin, setBaseCoin] = useState('');
  const [targetCoin, setTargetCoin] = useState('');
  const [currency, setCurrency] = useState([]);
  
  return (
      <div className='grid grid-cols-2 p-5'>
        <div className='col-auto'>
          <div className='p-1'>
            <label className='block mb-2 text-sm text-slate-600'>Base coin:</label>
            <input className='w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow' 
              type='text' 
              onChange={(event) => typingHandler(event, setBaseCoin)}/>
          </div>
          <div className='p-1' title='Use , to pass many target coins'>
            <label className='block mb-2 text-sm text-slate-600'>Target coin(s):</label>
            <input className='w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow' 
              type='text' 
              onChange={(event) => typingHandler(event, setTargetCoin)} />
          </div>
          <div className='p-1 text-center'>
            <button 
              className={`
                font-bold py-1 px-4 rounded 
                ${!baseCoin.trim() || !targetCoin.trim() 
                  ? 'bg-gray-400 cursor-not-allowed'   // Disabled state styles
                  : 'bg-blue-500 hover:bg-blue-700 text-white'  // Enabled state styles
                }
              `} 
              disabled={!baseCoin.trim() || !targetCoin.trim()}
              onClick={() => convert(baseCoin, targetCoin, setCurrency)}>Convert</button>
          </div>
        </div>
        <div className='col-auto'>
          <div className={`
            ${!baseCoin.trim() || !targetCoin.trim() 
              ? 'invisible'
              : ''
            }
            bg-blue-500 p-1 rounded`}>
            <label className='text-white'>Converting 1 {baseCoin} to:</label>
          </div>
          {currency.map((c:any) => (
          <div className='bg-gray-200 m-1 rounded'>
            <label>{c.exchange.name}: {c.exchange.price}</label>
          </div>
          ))}
        </div>
      </div>
    );
  }

export default CurrencyConverter;