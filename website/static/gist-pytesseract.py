import time
import os
import sys
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor
from typing import List, Union

import PIL
import pytesseract

img_dir = "/images"  # directory that contains the test images


def get_images(n_images: int) -> List[PIL.Image]:
    images = []
    for image_file in os.listdir(img_dir)[:n_images]:
        image = PIL.open(os.path.join(img_dir, image_file))
        images.append(image)
    return images


def get_chunks(images: List[PIL.Image], n_chunks: int) -> List[List[PIL.Image]]:

    return


def run_ocr(worker_pool: Union[ThreadPoolExecutor, ProcessPoolExecutor], n_images: int, n_chunks: int) -> None:
    # get list of PIL images
    images = get_images(n_images)

    # divide into chunks (optional)
    chunks = None
    if n_chunks:
        chunks = get_chunks(images, n_chunks)

    # submit work to the pool
    with worker_pool as pool:
        futures = pool.map(pytesseract.image_to_data, images)

    # wait for results
    futures.result()
    return


if __name__ == '__main__':
    start = time.time()

    num_workers = sys.argv[1]
    num_images = sys.argv[2]
    num_chunks = sys.argv[3]

    if sys.argv[4] == 'thread':
        executor = ThreadPoolExecutor(num_workers)
    elif sys.argv[4] == 'process':
        executor = ProcessPoolExecutor(num_workers)
    else:
        print(f'bad arg for executor type! {sys.argv[4]} is not "thread" or "process"')
        sys.exit(1)

    run_ocr(executor, num_images, num_chunks)
    print(f'{__file__} took {time.time() - start} seconds')
