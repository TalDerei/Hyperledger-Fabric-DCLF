import React, { useState } from 'react';
import './App.css';

function App() {

  const [statusMsg, setStatusMsg] = useState(0)

  const clicked = () => {
    setStatusMsg("Currently unimplemented.")
  }

  return (
    <div className="App">
      <h1>balance-transfer</h1>
      <div class="Buttons">
        <button onClick={clicked}>Register and enroll user</button>
        <button onClick={clicked}>Create channel</button>
        <button onClick={clicked}>Join channel</button>
        <button onClick={clicked}>Update anchor peers</button>
        <button onClick={clicked}>Install chaincode on target peers</button>
        <button onClick={clicked}>Instantiate chaincode on target peers</button>
        <button onClick={clicked}>Invoke transaction on chaincode on target peers</button>
        <button onClick={clicked}>Query on chaincode on target peers</button>
        <button onClick={clicked}>Get Block by BlockNumber</button>
        <button onClick={clicked}>Get Transaction by Transaction ID</button>
        <button onClick={clicked}>Get Block by Hash</button>
        <button onClick={clicked}>Query for Channel Information</button>
        <button onClick={clicked}>Query for Channel instantiated chaincodes</button>
        <button onClick={clicked}>Fetch all Installed/instantiated chaincodes</button>
        <button onClick={clicked}>Fetch channels</button>
      </div>
      {
        statusMsg == 0 ? <div></div> : <p>{statusMsg}</p>
      }
    </div>
  );
}

export default App;
