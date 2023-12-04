from aiohttp import ClientSession, FormData, ContentTypeError
from aiohttp.web import Response
from settings import converter_service_url


async def send_file(url, file_data, filename):
    async with ClientSession() as session:
        data = FormData()
        data.add_field('file', file_data, filename=filename)

        async with session.post(url, data=data) as response:
            try:
                return await response.json()
            except ContentTypeError:
                return {'errors': await response.text()}


async def send_request_to_convert_file(url, file_id, convert_to):
    async with ClientSession() as session:

        async with session.post(url, params={'file_id': file_id, 'convert_to': convert_to}) as response:
            return await response.json()


async def download_converted_file(url, file_id):
    async with ClientSession() as session:
        async with session.get(url, params={'file_id': file_id}) as response:
            return await response.content.read()


async def convert_file(request):
    data = await request.post()
    convert_to = data.get('convert_to', '')
    uploaded_file = data.get('file')
    upload_response = await send_file(f'{converter_service_url}/upload', uploaded_file.file.read(), uploaded_file.filename)
    if upload_response.get('errors'):
        return Response(upload_response, status=400)
    converted_file_id = await send_request_to_convert_file(
        f'{converter_service_url}/convert', upload_response.get('file_id'), convert_to
    )
    converted_file = await download_converted_file(
        f'{converter_service_url}/download', converted_file_id.get('new_file_id')
    )
    data = FormData()
    data.add_field('file', converted_file, filename=f'converted_file.{convert_to}')

    return Response(body=data(), headers={})
