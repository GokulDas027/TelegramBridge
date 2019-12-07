  # Step by Step instruction for Using this GitHub Action.
  
  ## Step 1
  * Telegram
    * Create Telegram Bot and Find it's token id from @BotFather bot in Telegram [find more](https://core.telegram.org/bots).
    * You cand direct the bot message to a Channel, Group or You.Just Find the Chat Id of You, Group or Channel. It's a number. [Stack Overflow](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id). better search specifically.
  ## Step 2
  * Save the Bot token and Chat id in the Repository Secrets(Settings/Secrets) as 2 different secrets.
    * name: chat  value:____chat id_____
    * name: token  value:____bot token____
  ## Step 3
  * Add Action on Repo
    * Click on Repo tab
    * Click "setup workflow yourself" on right top
    * copy paste the [main.yaml](https://github.com/GokulDas027/TelegramBridge/blob/master/main.yaml) in this repo to the opened main.yaml in your repo
 
 Yeah It's done. You will get notification from the bot in your channel/group/you.
 nb: add the bot as admin for channel and maybe for group too.
 
