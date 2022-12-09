import csv
import json
import logging
import os
import os.path
import shutil
import tempfile
import urllib.request
import zipfile

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s %(levelname)s [%(funcName)s] %(message)s",
)

log: logging.Logger = logging.getLogger(__name__)

DATA_PACKAGE_JSON_URL = "https://data.gov.ua/dataset/06779371-308f-42d7-895e-5a39833375f0/datapackage"

TMP_DIR: str = tempfile.mkdtemp()


def download_data_package_json() -> str:
    log.info("Downloading data package json")
    log.debug("Tmp directory path: %s", TMP_DIR)

    result_path: str = "{}/datapackage.json".format(TMP_DIR)
    (json_path, http_message) = urllib.request.urlretrieve(DATA_PACKAGE_JSON_URL,
                                                           result_path)
    log.debug("Download result path: %s", json_path)

    if json_path is None or json_path == "":
        log.error("message: %s", http_message)
        raise Exception("Error during downloading Data Package Json")

    log.info("Downloading data package json is finished")
    return result_path


def get_csv_links_from_json(json_file_path: str) -> list[str]:
    log.info("Will be parsed file by path")
    log.debug("File pathL %s", json_file_path)
    with open(json_file_path) as json_file:
        data = json.load(json_file)
        log.debug("parsed json: %s", data)

        resources: list[dict[str, str]] = data["resources"]
        log.debug("Resources: %s", resources)

        csv_file_links: list[str] = list(
            map((lambda resource: resource["path"]), resources))
        log.debug("csv links: %s", csv_file_links)
        log.info("Links were found, amount: %s", len(csv_file_links))
        return csv_file_links


def download_csv_archives(csv_links: list[str]) -> list[str]:
    log.info("Downloading csv archives")
    log.debug("Tmp directory path: %s", TMP_DIR)

    csv_archive_path_list: list[str] = []

    for link in csv_links:
        file_name: str = link.rsplit("/", 1)[-1]
        log.debug("Zip file name for download: %s", file_name)

        res_path: str = TMP_DIR + "/" + file_name
        log.debug("Zip file full path: %s", res_path)

        (csv_zip_path, http_message) = urllib.request.urlretrieve(link, res_path)
        log.debug("Download result path: %s", csv_zip_path)

        if csv_zip_path is None or csv_zip_path == "":
            log.error("message")
            log.error(http_message)
            raise Exception("Error during downloading CSV zip")

        csv_archive_path_list.append(csv_zip_path)

    log.debug("List of downloaded archives: %s", csv_archive_path_list)
    log.info("Downloading csv archives is finished")
    return csv_archive_path_list


def unpack_archives(csv_archives: list[str]) -> list[str]:
    log.info("Unpacking csv archives")
    log.debug(csv_archives)
    zip_folders: list[str] = []
    for zip_file_path in csv_archives:
        file_name: str = zip_file_path.rsplit("/", 1)[-1]
        log.debug("File name: %s", file_name)

        folder_name: str = file_name.rsplit(".", 1)[0]
        log.debug("Folder name: %s", folder_name)

        full_path: str = TMP_DIR + "/" + folder_name
        log.debug("Full path: %s", full_path)

        with zipfile.ZipFile(zip_file_path, 'r') as zip_file:
            zip_file.extractall(full_path)

        zip_folders.append(full_path)

    log.info("Unpacking csv archives is finished")
    log.debug(zip_folders)
    return zip_folders


def read_csv_files(folders: list[str], process_record_func) -> None:
    log.info("Reading CSV files")
    log.debug(folders)
    for csv_file_path in folders:
        all_files = os.listdir(csv_file_path)
        log.debug(all_files)

        def is_file(file_path, file):
            log.debug("Checking: %s, %s", file_path, file)
            return os.path.isfile(os.path.join(file_path, file))

        all_csv_files: list[str] = [file for file in all_files if
                                    is_file(csv_file_path, file)]
        log.debug(all_csv_files)

        for file_to_process in all_csv_files:
            log.debug(file_to_process)
            with open(file_to_process, 'r') as file:
                csv_reader: csv.DictReader = csv.DictReader(f=file, delimiter=";")
                for record in csv_reader:
                    dict_from_csv = dict(record)
                    normalized = normalize_dict_rec(dict_from_csv)
                    process_record_func(normalized)
    log.info("Reading CSV files is finished")


def normalize_dict_rec(record: dict[str, str]) -> dict[str, str]:
    new_dict: dict[str, str] = {}
    for key, value in record.items():
        new_dict[key.upper()] = value.upper()
    return new_dict


def process_record(record: dict[str, str]) -> None:
    log.info(record)


def delete_folder(folder_path: str):
    log.info("Tmp folder deleting")
    log.debug(folder_path)
    shutil.rmtree(folder_path, True)
    log.info("Folder removed")


if __name__ == "__main__":
    path = download_data_package_json()
    links = get_csv_links_from_json(path)
    csv_archives = download_csv_archives(links)
    unpacked_folders = unpack_archives(csv_archives)
    read_csv_files(unpacked_folders, process_record)
    delete_folder(TMP_DIR)
