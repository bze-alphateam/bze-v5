<!DOCTYPE html>

<html>
  <head>
    <title>Protocol Documentation</title>
    <meta charset="UTF-8">
    <link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Ubuntu:400,700,400italic"/>
    <style>
      body {
        width: 60em;
        margin: 1em auto;
        color: #222;
        font-family: "Ubuntu", sans-serif;
        padding-bottom: 4em;
      }

      h1 {
        font-weight: normal;
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
      }

      h2 {
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
        margin: 1.5em 0;
      }

      h3 {
        font-weight: normal;
        border-bottom: 1px solid #aaa;
        padding-bottom: 0.5ex;
      }

      a {
        text-decoration: none;
        color: #567e25;
      }

      table {
        width: 100%;
        font-size: 80%;
        border-collapse: collapse;
      }

      thead {
        font-weight: 700;
        background-color: #dcdcdc;
      }

      tbody tr:nth-child(even) {
        background-color: #fbfbfb;
      }

      td {
        border: 1px solid #ccc;
        padding: 0.5ex 2ex;
      }

      td p {
        text-indent: 1em;
        margin: 0;
      }

      td p:nth-child(1) {
        text-indent: 0;  
      }

       
      .field-table td:nth-child(1) {  
        width: 10em;
      }
      .field-table td:nth-child(2) {  
        width: 10em;
      }
      .field-table td:nth-child(3) {  
        width: 6em;
      }
      .field-table td:nth-child(4) {  
        width: auto;
      }

       
      .extension-table td:nth-child(1) {  
        width: 10em;
      }
      .extension-table td:nth-child(2) {  
        width: 10em;
      }
      .extension-table td:nth-child(3) {  
        width: 10em;
      }
      .extension-table td:nth-child(4) {  
        width: 5em;
      }
      .extension-table td:nth-child(5) {  
        width: auto;
      }

       
      .enum-table td:nth-child(1) {  
        width: 10em;
      }
      .enum-table td:nth-child(2) {  
        width: 10em;
      }
      .enum-table td:nth-child(3) {  
        width: auto;
      }

       
      .scalar-value-types-table tr {
        height: 3em;
      }

       
      #toc-container ul {
        list-style-type: none;
        padding-left: 1em;
        line-height: 180%;
        margin: 0;
      }
      #toc > li > a {
        font-weight: bold;
      }

       
      .file-heading {
        width: 100%;
        display: table;
        border-bottom: 1px solid #aaa;
        margin: 4em 0 1.5em 0;
      }
      .file-heading h2 {
        border: none;
        display: table-cell;
      }
      .file-heading a {
        text-align: right;
        display: table-cell;
      }

       
      .badge {
        width: 1.6em;
        height: 1.6em;
        display: inline-block;

        line-height: 1.6em;
        text-align: center;
        font-weight: bold;
        font-size: 60%;

        color: #89ba48;
        background-color: #dff0c8;

        margin: 0.5ex 1em 0.5ex -1em;
        border: 1px solid #fbfbfb;
        border-radius: 1ex;
      }
    </style>

    
    <link rel="stylesheet" type="text/css" href="stylesheet.css"/>
  </head>

  <body>

    <h1 id="title">Protocol Documentation</h1>

    <h2>Table of Contents</h2>

    <div id="toc-container">
      <ul id="toc">
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fevents.proto">bze/rewards/v1/events.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardClaimEvent"><span class="badge">M</span>StakingRewardClaimEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardCreateEvent"><span class="badge">M</span>StakingRewardCreateEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardDistributionEvent"><span class="badge">M</span>StakingRewardDistributionEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardExitEvent"><span class="badge">M</span>StakingRewardExitEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardFinishEvent"><span class="badge">M</span>StakingRewardFinishEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardJoinEvent"><span class="badge">M</span>StakingRewardJoinEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardUpdateEvent"><span class="badge">M</span>StakingRewardUpdateEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardActivationEvent"><span class="badge">M</span>TradingRewardActivationEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardCreateEvent"><span class="badge">M</span>TradingRewardCreateEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardDistributionEvent"><span class="badge">M</span>TradingRewardDistributionEvent</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardExpireEvent"><span class="badge">M</span>TradingRewardExpireEvent</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fparams.proto">bze/rewards/v1/params.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.Params"><span class="badge">M</span>Params</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fstaking_reward.proto">bze/rewards/v1/staking_reward.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.StakingReward"><span class="badge">M</span>StakingReward</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2ftrading_reward.proto">bze/rewards/v1/trading_reward.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.MarketIdTradingRewardId"><span class="badge">M</span>MarketIdTradingRewardId</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingReward"><span class="badge">M</span>TradingReward</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardCandidate"><span class="badge">M</span>TradingRewardCandidate</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardExpiration"><span class="badge">M</span>TradingRewardExpiration</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardLeaderboard"><span class="badge">M</span>TradingRewardLeaderboard</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.TradingRewardLeaderboardEntry"><span class="badge">M</span>TradingRewardLeaderboardEntry</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fstaking_reward_participant.proto">bze/rewards/v1/staking_reward_participant.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.PendingUnlockParticipant"><span class="badge">M</span>PendingUnlockParticipant</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.StakingRewardParticipant"><span class="badge">M</span>StakingRewardParticipant</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fgenesis.proto">bze/rewards/v1/genesis.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.GenesisState"><span class="badge">M</span>GenesisState</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fgov.proto">bze/rewards/v1/gov.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.ActivateTradingRewardProposal"><span class="badge">M</span>ActivateTradingRewardProposal</a>
                </li>
              
              
              
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2fquery.proto">bze/rewards/v1/query.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllPendingUnlockParticipantRequest"><span class="badge">M</span>QueryAllPendingUnlockParticipantRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllPendingUnlockParticipantResponse"><span class="badge">M</span>QueryAllPendingUnlockParticipantResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllStakingRewardParticipantRequest"><span class="badge">M</span>QueryAllStakingRewardParticipantRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllStakingRewardParticipantResponse"><span class="badge">M</span>QueryAllStakingRewardParticipantResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllStakingRewardRequest"><span class="badge">M</span>QueryAllStakingRewardRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllStakingRewardResponse"><span class="badge">M</span>QueryAllStakingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllTradingRewardRequest"><span class="badge">M</span>QueryAllTradingRewardRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryAllTradingRewardResponse"><span class="badge">M</span>QueryAllTradingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerRequest"><span class="badge">M</span>QueryGetMarketIdTradingRewardIdHandlerRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerResponse"><span class="badge">M</span>QueryGetMarketIdTradingRewardIdHandlerResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetStakingRewardParticipantRequest"><span class="badge">M</span>QueryGetStakingRewardParticipantRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetStakingRewardParticipantResponse"><span class="badge">M</span>QueryGetStakingRewardParticipantResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetStakingRewardRequest"><span class="badge">M</span>QueryGetStakingRewardRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetStakingRewardResponse"><span class="badge">M</span>QueryGetStakingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetTradingRewardLeaderboardRequest"><span class="badge">M</span>QueryGetTradingRewardLeaderboardRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetTradingRewardLeaderboardResponse"><span class="badge">M</span>QueryGetTradingRewardLeaderboardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetTradingRewardRequest"><span class="badge">M</span>QueryGetTradingRewardRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryGetTradingRewardResponse"><span class="badge">M</span>QueryGetTradingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryParamsRequest"><span class="badge">M</span>QueryParamsRequest</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.QueryParamsResponse"><span class="badge">M</span>QueryParamsResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#bze.rewards.v1.Query"><span class="badge">S</span>Query</a>
                </li>
              
            </ul>
          </li>
        
          
          <li>
            <a href="#bze%2frewards%2fv1%2ftx.proto">bze/rewards/v1/tx.proto</a>
            <ul>
              
                <li>
                  <a href="#bze.rewards.v1.MsgClaimStakingRewards"><span class="badge">M</span>MsgClaimStakingRewards</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgClaimStakingRewardsResponse"><span class="badge">M</span>MsgClaimStakingRewardsResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgCreateStakingReward"><span class="badge">M</span>MsgCreateStakingReward</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgCreateStakingRewardResponse"><span class="badge">M</span>MsgCreateStakingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgCreateTradingReward"><span class="badge">M</span>MsgCreateTradingReward</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgCreateTradingRewardResponse"><span class="badge">M</span>MsgCreateTradingRewardResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgDistributeStakingRewards"><span class="badge">M</span>MsgDistributeStakingRewards</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgDistributeStakingRewardsResponse"><span class="badge">M</span>MsgDistributeStakingRewardsResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgExitStaking"><span class="badge">M</span>MsgExitStaking</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgExitStakingResponse"><span class="badge">M</span>MsgExitStakingResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgJoinStaking"><span class="badge">M</span>MsgJoinStaking</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgJoinStakingResponse"><span class="badge">M</span>MsgJoinStakingResponse</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgUpdateStakingReward"><span class="badge">M</span>MsgUpdateStakingReward</a>
                </li>
              
                <li>
                  <a href="#bze.rewards.v1.MsgUpdateStakingRewardResponse"><span class="badge">M</span>MsgUpdateStakingRewardResponse</a>
                </li>
              
              
              
              
                <li>
                  <a href="#bze.rewards.v1.Msg"><span class="badge">S</span>Msg</a>
                </li>
              
            </ul>
          </li>
        
        <li><a href="#scalar-value-types">Scalar Value Types</a></li>
      </ul>
    </div>

    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/events.proto">bze/rewards/v1/events.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.StakingRewardClaimEvent">StakingRewardClaimEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardCreateEvent">StakingRewardCreateEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staking_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>min_stake</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>lock</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardDistributionEvent">StakingRewardDistributionEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardExitEvent">StakingRewardExitEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardFinishEvent">StakingRewardFinishEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardJoinEvent">StakingRewardJoinEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardUpdateEvent">StakingRewardUpdateEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardActivationEvent">TradingRewardActivationEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardCreateEvent">TradingRewardCreateEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the amount paid as prize for each slot </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the denom paid as prize </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>market_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>slots</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardDistributionEvent">TradingRewardDistributionEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>winners</td>
                  <td><a href="#string">string</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardExpireEvent">TradingRewardExpireEvent</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/params.proto">bze/rewards/v1/params.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.Params">Params</h3>
        <p>Params defines the parameters for the module.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>createStakingRewardFee</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>createTradingRewardFee</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/staking_reward.proto">bze/rewards/v1/staking_reward.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.StakingReward">StakingReward</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staking_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>payouts</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>min_stake</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>lock</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staked_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>T </p></td>
                </tr>
              
                <tr>
                  <td>distributed_stake</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>S </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/trading_reward.proto">bze/rewards/v1/trading_reward.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.MarketIdTradingRewardId">MarketIdTradingRewardId</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>market_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingReward">TradingReward</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>market_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>slots</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>expire_at</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardCandidate">TradingRewardCandidate</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardExpiration">TradingRewardExpiration</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>expire_at</td>
                  <td><a href="#uint32">uint32</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardLeaderboard">TradingRewardLeaderboard</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.TradingRewardLeaderboardEntry">TradingRewardLeaderboardEntry</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.TradingRewardLeaderboardEntry">TradingRewardLeaderboardEntry</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>created_at</td>
                  <td><a href="#int64">int64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/staking_reward_participant.proto">bze/rewards/v1/staking_reward_participant.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.PendingUnlockParticipant">PendingUnlockParticipant</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>index</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.StakingRewardParticipant">StakingRewardParticipant</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>stake[address] </p></td>
                </tr>
              
                <tr>
                  <td>joined_at</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>S0[address] </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/genesis.proto">bze/rewards/v1/genesis.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.GenesisState">GenesisState</h3>
        <p>GenesisState defines the rewards module's genesis state.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
                  <td><a href="#bze.rewards.v1.Params">Params</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staking_reward_list</td>
                  <td><a href="#bze.rewards.v1.StakingReward">StakingReward</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staking_rewards_counter</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>trading_rewards_counter</td>
                  <td><a href="#uint64">uint64</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>active_trading_reward_list</td>
                  <td><a href="#bze.rewards.v1.TradingReward">TradingReward</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pending_trading_reward_list</td>
                  <td><a href="#bze.rewards.v1.TradingReward">TradingReward</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>staking_reward_participant_list</td>
                  <td><a href="#bze.rewards.v1.StakingRewardParticipant">StakingRewardParticipant</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pending_unlock_participant_list</td>
                  <td><a href="#bze.rewards.v1.PendingUnlockParticipant">PendingUnlockParticipant</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>trading_reward_leaderboard_list</td>
                  <td><a href="#bze.rewards.v1.TradingRewardLeaderboard">TradingRewardLeaderboard</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>trading_reward_candidate_list</td>
                  <td><a href="#bze.rewards.v1.TradingRewardCandidate">TradingRewardCandidate</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>market_id_trading_reward_id_list</td>
                  <td><a href="#bze.rewards.v1.MarketIdTradingRewardId">MarketIdTradingRewardId</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pending_trading_reward_expiration_list</td>
                  <td><a href="#bze.rewards.v1.TradingRewardExpiration">TradingRewardExpiration</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>active_trading_reward_expiration_list</td>
                  <td><a href="#bze.rewards.v1.TradingRewardExpiration">TradingRewardExpiration</a></td>
                  <td>repeated</td>
                  <td><p>this line is used by starport scaffolding # genesis/proto/state </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/gov.proto">bze/rewards/v1/gov.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.ActivateTradingRewardProposal">ActivateTradingRewardProposal</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>title</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>description</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/query.proto">bze/rewards/v1/query.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.QueryAllPendingUnlockParticipantRequest">QueryAllPendingUnlockParticipantRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageRequest">cosmos.base.query.v1beta1.PageRequest</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllPendingUnlockParticipantResponse">QueryAllPendingUnlockParticipantResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.PendingUnlockParticipant">PendingUnlockParticipant</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageResponse">cosmos.base.query.v1beta1.PageResponse</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllStakingRewardParticipantRequest">QueryAllStakingRewardParticipantRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageRequest">cosmos.base.query.v1beta1.PageRequest</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllStakingRewardParticipantResponse">QueryAllStakingRewardParticipantResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.StakingRewardParticipant">StakingRewardParticipant</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageResponse">cosmos.base.query.v1beta1.PageResponse</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllStakingRewardRequest">QueryAllStakingRewardRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageRequest">cosmos.base.query.v1beta1.PageRequest</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllStakingRewardResponse">QueryAllStakingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.StakingReward">StakingReward</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageResponse">cosmos.base.query.v1beta1.PageResponse</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllTradingRewardRequest">QueryAllTradingRewardRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>state</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageRequest">cosmos.base.query.v1beta1.PageRequest</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryAllTradingRewardResponse">QueryAllTradingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.TradingReward">TradingReward</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageResponse">cosmos.base.query.v1beta1.PageResponse</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerRequest">QueryGetMarketIdTradingRewardIdHandlerRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>market_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerResponse">QueryGetMarketIdTradingRewardIdHandlerResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>market_id_reward_id</td>
                  <td><a href="#bze.rewards.v1.MarketIdTradingRewardId">MarketIdTradingRewardId</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetStakingRewardParticipantRequest">QueryGetStakingRewardParticipantRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>address</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageRequest">cosmos.base.query.v1beta1.PageRequest</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetStakingRewardParticipantResponse">QueryGetStakingRewardParticipantResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>list</td>
                  <td><a href="#bze.rewards.v1.StakingRewardParticipant">StakingRewardParticipant</a></td>
                  <td>repeated</td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>pagination</td>
                  <td><a href="#cosmos.base.query.v1beta1.PageResponse">cosmos.base.query.v1beta1.PageResponse</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetStakingRewardRequest">QueryGetStakingRewardRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetStakingRewardResponse">QueryGetStakingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>staking_reward</td>
                  <td><a href="#bze.rewards.v1.StakingReward">StakingReward</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetTradingRewardLeaderboardRequest">QueryGetTradingRewardLeaderboardRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetTradingRewardLeaderboardResponse">QueryGetTradingRewardLeaderboardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>leaderboard</td>
                  <td><a href="#bze.rewards.v1.TradingRewardLeaderboard">TradingRewardLeaderboard</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetTradingRewardRequest">QueryGetTradingRewardRequest</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryGetTradingRewardResponse">QueryGetTradingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>trading_reward</td>
                  <td><a href="#bze.rewards.v1.TradingReward">TradingReward</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.QueryParamsRequest">QueryParamsRequest</h3>
        <p>QueryParamsRequest is request type for the Query/Params RPC method.</p>

        

        
      
        <h3 id="bze.rewards.v1.QueryParamsResponse">QueryParamsResponse</h3>
        <p>QueryParamsResponse is response type for the Query/Params RPC method.</p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>params</td>
                  <td><a href="#bze.rewards.v1.Params">Params</a></td>
                  <td></td>
                  <td><p>params holds all the parameters of this module. </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      

      

      

      
        <h3 id="bze.rewards.v1.Query">Query</h3>
        <p>Query defines the gRPC querier service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>Params</td>
                <td><a href="#bze.rewards.v1.QueryParamsRequest">QueryParamsRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryParamsResponse">QueryParamsResponse</a></td>
                <td><p>Parameters queries the parameters of the module.</p></td>
              </tr>
            
              <tr>
                <td>StakingReward</td>
                <td><a href="#bze.rewards.v1.QueryGetStakingRewardRequest">QueryGetStakingRewardRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryGetStakingRewardResponse">QueryGetStakingRewardResponse</a></td>
                <td><p>Queries a StakingReward by index.</p></td>
              </tr>
            
              <tr>
                <td>StakingRewardAll</td>
                <td><a href="#bze.rewards.v1.QueryAllStakingRewardRequest">QueryAllStakingRewardRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryAllStakingRewardResponse">QueryAllStakingRewardResponse</a></td>
                <td><p>Queries a list of StakingReward items.</p></td>
              </tr>
            
              <tr>
                <td>TradingReward</td>
                <td><a href="#bze.rewards.v1.QueryGetTradingRewardRequest">QueryGetTradingRewardRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryGetTradingRewardResponse">QueryGetTradingRewardResponse</a></td>
                <td><p>Queries a TradingReward by index.</p></td>
              </tr>
            
              <tr>
                <td>TradingRewardAll</td>
                <td><a href="#bze.rewards.v1.QueryAllTradingRewardRequest">QueryAllTradingRewardRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryAllTradingRewardResponse">QueryAllTradingRewardResponse</a></td>
                <td><p>Queries a list of TradingReward items.</p></td>
              </tr>
            
              <tr>
                <td>StakingRewardParticipant</td>
                <td><a href="#bze.rewards.v1.QueryGetStakingRewardParticipantRequest">QueryGetStakingRewardParticipantRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryGetStakingRewardParticipantResponse">QueryGetStakingRewardParticipantResponse</a></td>
                <td><p>Queries a StakingRewardParticipant by index.</p></td>
              </tr>
            
              <tr>
                <td>StakingRewardParticipantAll</td>
                <td><a href="#bze.rewards.v1.QueryAllStakingRewardParticipantRequest">QueryAllStakingRewardParticipantRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryAllStakingRewardParticipantResponse">QueryAllStakingRewardParticipantResponse</a></td>
                <td><p>Queries a list of StakingRewardParticipant items.</p></td>
              </tr>
            
              <tr>
                <td>GetTradingRewardLeaderboardHandler</td>
                <td><a href="#bze.rewards.v1.QueryGetTradingRewardLeaderboardRequest">QueryGetTradingRewardLeaderboardRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryGetTradingRewardLeaderboardResponse">QueryGetTradingRewardLeaderboardResponse</a></td>
                <td><p>Queries a list of GetTradingRewardLeaderboard items.</p></td>
              </tr>
            
              <tr>
                <td>GetMarketIdTradingRewardIdHandler</td>
                <td><a href="#bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerRequest">QueryGetMarketIdTradingRewardIdHandlerRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryGetMarketIdTradingRewardIdHandlerResponse">QueryGetMarketIdTradingRewardIdHandlerResponse</a></td>
                <td><p>Queries a list of GetMarketIdTradingRewardIdHandler items.</p></td>
              </tr>
            
              <tr>
                <td>AllPendingUnlockParticipant</td>
                <td><a href="#bze.rewards.v1.QueryAllPendingUnlockParticipantRequest">QueryAllPendingUnlockParticipantRequest</a></td>
                <td><a href="#bze.rewards.v1.QueryAllPendingUnlockParticipantResponse">QueryAllPendingUnlockParticipantResponse</a></td>
                <td><p>Queries a list of AllPendingUnlockParticipant items.</p></td>
              </tr>
            
          </tbody>
        </table>

        
          
          
          <h4>Methods with HTTP bindings</h4>
          <table>
            <thead>
              <tr>
                <td>Method Name</td>
                <td>Method</td>
                <td>Pattern</td>
                <td>Body</td>
              </tr>
            </thead>
            <tbody>
            
              
              
              <tr>
                <td>Params</td>
                <td>GET</td>
                <td>/bze/rewards/v1/params</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>StakingReward</td>
                <td>GET</td>
                <td>/bze/rewards/v1/staking_reward/{reward_id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>StakingRewardAll</td>
                <td>GET</td>
                <td>/bze/rewards/v1/staking_reward</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>TradingReward</td>
                <td>GET</td>
                <td>/bze/rewards/v1/trading_reward/{reward_id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>TradingRewardAll</td>
                <td>GET</td>
                <td>/bze/rewards/v1/trading_reward/{state}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>StakingRewardParticipant</td>
                <td>GET</td>
                <td>/bze/rewards/v1/staking_reward_participant/{address}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>StakingRewardParticipantAll</td>
                <td>GET</td>
                <td>/bze/rewards/v1/staking_reward_participants</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>GetTradingRewardLeaderboardHandler</td>
                <td>GET</td>
                <td>/bze/rewards/v1/trading_reward_leaderboard/{reward_id}</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>GetMarketIdTradingRewardIdHandler</td>
                <td>GET</td>
                <td>/bze/rewards/v1/market_id_trading_reward_id</td>
                <td></td>
              </tr>
              
            
              
              
              <tr>
                <td>AllPendingUnlockParticipant</td>
                <td>GET</td>
                <td>/bze/rewards/v1/all_pending_unlock_participant</td>
                <td></td>
              </tr>
              
            
            </tbody>
          </table>
          
        
    
      
      <div class="file-heading">
        <h2 id="bze/rewards/v1/tx.proto">bze/rewards/v1/tx.proto</h2><a href="#title">Top</a>
      </div>
      <p></p>

      
        <h3 id="bze.rewards.v1.MsgClaimStakingRewards">MsgClaimStakingRewards</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>rewardId</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgClaimStakingRewardsResponse">MsgClaimStakingRewardsResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgCreateStakingReward">MsgCreateStakingReward</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>msg creator </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the amount paid as prize for each epoch (duration) </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the denom paid as prize </p></td>
                </tr>
              
                <tr>
                  <td>staking_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the denom a user has to stake in order to qualify </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the number of days the rewards are paid </p></td>
                </tr>
              
                <tr>
                  <td>min_stake</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the minimum amount of staking denom a user has to stake in order to qualify </p></td>
                </tr>
              
                <tr>
                  <td>lock</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the number of days the funds are locked upon exiting stake </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgCreateStakingRewardResponse">MsgCreateStakingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgCreateTradingReward">MsgCreateTradingReward</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>prize_amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the amount paid as prize for each slot </p></td>
                </tr>
              
                <tr>
                  <td>prize_denom</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the denom paid as prize </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>market_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>slots</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgCreateTradingRewardResponse">MsgCreateTradingRewardResponse</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgDistributeStakingRewards">MsgDistributeStakingRewards</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>rewardId</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgDistributeStakingRewardsResponse">MsgDistributeStakingRewardsResponse</h3>
        <p></p>

        

        
      
        <h3 id="bze.rewards.v1.MsgExitStaking">MsgExitStaking</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>rewardId</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgExitStakingResponse">MsgExitStakingResponse</h3>
        <p></p>

        

        
      
        <h3 id="bze.rewards.v1.MsgJoinStaking">MsgJoinStaking</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>amount</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgJoinStakingResponse">MsgJoinStakingResponse</h3>
        <p></p>

        

        
      
        <h3 id="bze.rewards.v1.MsgUpdateStakingReward">MsgUpdateStakingReward</h3>
        <p></p>

        
          <table class="field-table">
            <thead>
              <tr><td>Field</td><td>Type</td><td>Label</td><td>Description</td></tr>
            </thead>
            <tbody>
              
                <tr>
                  <td>creator</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>reward_id</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p> </p></td>
                </tr>
              
                <tr>
                  <td>duration</td>
                  <td><a href="#string">string</a></td>
                  <td></td>
                  <td><p>the number of days the rewards are paid </p></td>
                </tr>
              
            </tbody>
          </table>

          

        
      
        <h3 id="bze.rewards.v1.MsgUpdateStakingRewardResponse">MsgUpdateStakingRewardResponse</h3>
        <p></p>

        

        
      

      

      

      
        <h3 id="bze.rewards.v1.Msg">Msg</h3>
        <p>Msg defines the Msg service.</p>
        <table class="enum-table">
          <thead>
            <tr><td>Method Name</td><td>Request Type</td><td>Response Type</td><td>Description</td></tr>
          </thead>
          <tbody>
            
              <tr>
                <td>CreateStakingReward</td>
                <td><a href="#bze.rewards.v1.MsgCreateStakingReward">MsgCreateStakingReward</a></td>
                <td><a href="#bze.rewards.v1.MsgCreateStakingRewardResponse">MsgCreateStakingRewardResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>UpdateStakingReward</td>
                <td><a href="#bze.rewards.v1.MsgUpdateStakingReward">MsgUpdateStakingReward</a></td>
                <td><a href="#bze.rewards.v1.MsgUpdateStakingRewardResponse">MsgUpdateStakingRewardResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>CreateTradingReward</td>
                <td><a href="#bze.rewards.v1.MsgCreateTradingReward">MsgCreateTradingReward</a></td>
                <td><a href="#bze.rewards.v1.MsgCreateTradingRewardResponse">MsgCreateTradingRewardResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>JoinStaking</td>
                <td><a href="#bze.rewards.v1.MsgJoinStaking">MsgJoinStaking</a></td>
                <td><a href="#bze.rewards.v1.MsgJoinStakingResponse">MsgJoinStakingResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>ExitStaking</td>
                <td><a href="#bze.rewards.v1.MsgExitStaking">MsgExitStaking</a></td>
                <td><a href="#bze.rewards.v1.MsgExitStakingResponse">MsgExitStakingResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>ClaimStakingRewards</td>
                <td><a href="#bze.rewards.v1.MsgClaimStakingRewards">MsgClaimStakingRewards</a></td>
                <td><a href="#bze.rewards.v1.MsgClaimStakingRewardsResponse">MsgClaimStakingRewardsResponse</a></td>
                <td><p></p></td>
              </tr>
            
              <tr>
                <td>DistributeStakingRewards</td>
                <td><a href="#bze.rewards.v1.MsgDistributeStakingRewards">MsgDistributeStakingRewards</a></td>
                <td><a href="#bze.rewards.v1.MsgDistributeStakingRewardsResponse">MsgDistributeStakingRewardsResponse</a></td>
                <td><p>this line is used by starport scaffolding # proto/tx/rpc</p></td>
              </tr>
            
          </tbody>
        </table>

        
    

    <h2 id="scalar-value-types">Scalar Value Types</h2>
    <table class="scalar-value-types-table">
      <thead>
        <tr><td>.proto Type</td><td>Notes</td><td>C++</td><td>Java</td><td>Python</td><td>Go</td><td>C#</td><td>PHP</td><td>Ruby</td></tr>
      </thead>
      <tbody>
        
          <tr id="double">
            <td>double</td>
            <td></td>
            <td>double</td>
            <td>double</td>
            <td>float</td>
            <td>float64</td>
            <td>double</td>
            <td>float</td>
            <td>Float</td>
          </tr>
        
          <tr id="float">
            <td>float</td>
            <td></td>
            <td>float</td>
            <td>float</td>
            <td>float</td>
            <td>float32</td>
            <td>float</td>
            <td>float</td>
            <td>Float</td>
          </tr>
        
          <tr id="int32">
            <td>int32</td>
            <td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="int64">
            <td>int64</td>
            <td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="uint32">
            <td>uint32</td>
            <td>Uses variable-length encoding.</td>
            <td>uint32</td>
            <td>int</td>
            <td>int/long</td>
            <td>uint32</td>
            <td>uint</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="uint64">
            <td>uint64</td>
            <td>Uses variable-length encoding.</td>
            <td>uint64</td>
            <td>long</td>
            <td>int/long</td>
            <td>uint64</td>
            <td>ulong</td>
            <td>integer/string</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sint32">
            <td>sint32</td>
            <td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sint64">
            <td>sint64</td>
            <td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="fixed32">
            <td>fixed32</td>
            <td>Always four bytes. More efficient than uint32 if values are often greater than 2^28.</td>
            <td>uint32</td>
            <td>int</td>
            <td>int</td>
            <td>uint32</td>
            <td>uint</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="fixed64">
            <td>fixed64</td>
            <td>Always eight bytes. More efficient than uint64 if values are often greater than 2^56.</td>
            <td>uint64</td>
            <td>long</td>
            <td>int/long</td>
            <td>uint64</td>
            <td>ulong</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="sfixed32">
            <td>sfixed32</td>
            <td>Always four bytes.</td>
            <td>int32</td>
            <td>int</td>
            <td>int</td>
            <td>int32</td>
            <td>int</td>
            <td>integer</td>
            <td>Bignum or Fixnum (as required)</td>
          </tr>
        
          <tr id="sfixed64">
            <td>sfixed64</td>
            <td>Always eight bytes.</td>
            <td>int64</td>
            <td>long</td>
            <td>int/long</td>
            <td>int64</td>
            <td>long</td>
            <td>integer/string</td>
            <td>Bignum</td>
          </tr>
        
          <tr id="bool">
            <td>bool</td>
            <td></td>
            <td>bool</td>
            <td>boolean</td>
            <td>boolean</td>
            <td>bool</td>
            <td>bool</td>
            <td>boolean</td>
            <td>TrueClass/FalseClass</td>
          </tr>
        
          <tr id="string">
            <td>string</td>
            <td>A string must always contain UTF-8 encoded or 7-bit ASCII text.</td>
            <td>string</td>
            <td>String</td>
            <td>str/unicode</td>
            <td>string</td>
            <td>string</td>
            <td>string</td>
            <td>String (UTF-8)</td>
          </tr>
        
          <tr id="bytes">
            <td>bytes</td>
            <td>May contain any arbitrary sequence of bytes.</td>
            <td>string</td>
            <td>ByteString</td>
            <td>str</td>
            <td>[]byte</td>
            <td>ByteString</td>
            <td>string</td>
            <td>String (ASCII-8BIT)</td>
          </tr>
        
      </tbody>
    </table>
  </body>
</html>

