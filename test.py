import aiohttp
import asyncio

conn = aiohttp.TCPConnector(limit=None, ttl_dns_cache=300)
session = aiohttp.ClientSession(connector=conn)

async def get_async(i):
    url = 'http://localhost:3000/request-id'
    async with session.post(url, ssl=False, headers={'X-Request-ID': i}) as response:
        obj = await response.json()
        print(obj['RequestID'], i)
        assert obj['RequestID'] == i

async def main():
    print("Testing")
    N = 10
    await asyncio.gather(*[get_async(str(i)) for i in range(N)])


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())