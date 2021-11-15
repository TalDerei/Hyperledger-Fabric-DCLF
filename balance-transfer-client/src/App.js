import React, { useState } from 'react';
import './App.css';

function App() {

  const [statusMsg, setStatusMsg] = useState('')
  const [networkParams, setNetworkParams] = useState('')
  const [queryParams, setQueryParams] = useState('')

  const clicked = () => {
    setStatusMsg("Currently unimplemented.")

  }

  const handleNetworkParamsChange = (event) => {
    setNetworkParams(event.target.value)
  }

  const handleQueryParamsChange = (event) => {
    setQueryParams(event.target.value)
  }

  const parseParams = (params) => {
    return params.split(',')
  }

  return (
    <div className="App">
      <h1>balance-transfer</h1>
      <div class="network-ops">
        <h2>Network Operations</h2>
        <input type="text" id="network-params" name="network-params" placeholder="parameters" 
          value={networkParams} onChange={handleNetworkParamsChange}/>
        <div class="buttons">
          <button onClick={clicked}>Register and enroll user</button>
          <button onClick={clicked}>Create channel</button>
          <button onClick={clicked}>Join channel</button>
          <button onClick={clicked}>Update anchor peers</button>
          <button onClick={clicked}>Install chaincode on target peers</button>
          <button onClick={clicked}>Instantiate chaincode on target peers</button>
          <button onClick={clicked}>Invoke transaction on chaincode on target peers</button>
        </div>
      </div>
      <div class="query">
        <h2>Query</h2>
        <input type="text" id="query-params" name="query-params" placeholder="parameters"
          value={queryParams} onChange={handleQueryParamsChange}/>
        <div class="buttons">
          <button onClick={clicked}>Query on chaincode on target peers</button>
          <button onClick={clicked}>Get Block by BlockNumber</button>
          <button onClick={clicked}>Get Transaction by Transaction ID</button>
          <button onClick={clicked}>Get Block by Hash</button>
          <button onClick={clicked}>Query for Channel Information</button>
          <button onClick={clicked}>Query for Channel instantiated chaincodes</button>
          <button onClick={clicked}>Fetch all Installed/instantiated chaincodes</button>
          <button onClick={clicked}>Fetch channels</button>
        </div>
      </div>
      <p style={{ "color": "red" }}>{statusMsg}</p>
      <div class="help">
        <h3>Help</h3>
        <p>Enter comma-separated params into text field (space and case sensitive)</p>
        <p>Enter params in order of use in balance-transfer endpoints</p>
        <p>Click on command to call it after entering params</p>
      </div>
    </div>
  );
}

export default App;
