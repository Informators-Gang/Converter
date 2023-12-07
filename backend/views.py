from aiohttp import ClientSession, FormData
from settings import converter_service_url


async def send_file(url, file_path):
    async with ClientSession() as session:
        data = FormData()
        data.add_field('file', open(file_path, 'rb'))

        async with session.post(url, data=data) as response:
            return await response.json()


async def send_request_to_convert_file(url, file_id, convert_to):
    async with ClientSession() as session:

        async with session.post(url, params={'file_id': file_id, 'convert_to': convert_to}) as response:
            return await response.json()


async def convert_file(request):
    data = await request.json()
    convert_to = data.get('convert_to', '')
    reader = await request.multipart()
    field = await reader.next()
    with open(field.filename, 'wb') as f:
        while True:
            chunk = await field.read_chunk()
            if not chunk:
                break
            f.write(chunk)
    file_id = await send_file(f'{converter_service_url}/upload', field.filename)
    converted_file_id = await send_request_to_convert_file(
        f'{converter_service_url}/convert', file_id.get('file_id'), convert_to
    )
