# StockSim

Simulates buying and selling stocks for real prices.

It fetches prices from this API: https://api-ninjas.com/api/stockprice

**To run it you will need an ApiNinjas API key**

Your portfolio is saved in a Player.json file to the directory from which you called the tool

## Command examples

### Buy
Buys one Apple stock
```powershell
StocksSim buy -t "AAPL" -n 1
```

### Sell
Sells one Apple stock
```powershell
StocksSim sell -t "AAPL" -n 1
```

### Portfolio
Prints your current portfolio
```powershell
StocksSim p
```
<img width="621" height="228" alt="portfolio_screenshot" src="https://github.com/user-attachments/assets/fbdeff33-ec69-4b2a-a872-0a9fe762c8d7" />

### Check a certain stock
```powershell
StocksSim check "AAPL"
```
<img width="464" height="208" alt="check_screenshot" src="https://github.com/user-attachments/assets/04111a45-2173-4225-a3b6-c3640c094b25" />
