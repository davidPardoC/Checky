# Check in and Check out track system.

The purpose of this project is to buiild a system to track check in and check out of employees in a restaurant or can be aplied in any businness if fits your needs. Also other of the purposes is to learn `go` and `gin framework` to build web apps. So backend is going to be builded with `go` and frontend will be made with `react`.

## Local Development

Recomended to use a live reload tool like [air](https://github.com/cosmtrek/air) to improve development experience.

1. Create a .env file.
   ```
   cp .env.example .env
   ```
2. Database:
   ```
   docker-compose -f docker-conpose.local.yml up -d
   ```
3. Development server: justo run

   ```
   air -c .air-unix.toml if youre developing in Mac/Linux , or
   air -c .air-windows.toml if you are developing in windows.
   ```

TO DO:

- Serve Static files.
- Ready to go setup.
