#!/bin/bash

echo "Starting database seeding process..."

mongosh <<EOF
use goparcel;

db.createCollection("user_types");
db.createCollection("locations");
db.createCollection("users");
db.createCollection("admin");

db.user_types.insertMany([
    {
        _id: ObjectId("67b54b11ba1542ee89a0f144"),
        name: "warehouse_operator",
        description: "Operator gudang yang menangani pengiriman dan penerimaan barang",
    },
    {
        _id: ObjectId("67b54b11ba1542ee89a0f145"),
        name: "depot_operator",
        description: "Operator depot yang bertugas mengelola distribusi barang",
    },
    {
        _id: ObjectId("67b54b11ba1542ee89a0f146"),
        name: "carrier",
        description: "Operator transportasi yang mengangkut barang ke tujuan",
    },
    {
        _id: ObjectId("67b54b11ba1542ee89a0f147"),
        name: "courier",
        description: "Kurir yang bertanggung jawab mengantarkan paket ke pelanggan",
    },
    {
        _id: ObjectId("67c000000000000000000002"),
        name: "admin",
        description: "Administrator dengan akses penuh ke seluruh sistem",
    }
]);

db.admin.insertOne({
    "_id": ObjectId("67b54b11ba1542ee89a0f148"),
    "user_id": ObjectId("67b54b11ba1542ee89a0f149"),
    "name": "Super Admin",
    "email": "super@admin.com",
});

db.users.insertOne({
    "_id": ObjectId("67b54b11ba1542ee89a0f149"),
    "model_id": ObjectId("67b54b11ba1542ee89a0f148"),
    "email": "super@admin.com",
    "entity": "admin",
    password: "\$2a\$12\$zesDtHwCbtHyPlAZ24cPjOjKaCHkeh7q3WlwofXYV8TEFM7b2Hapi", // hashed password
    user_type_id: ObjectId("67c000000000000000000002"),
});

db.locations.insertMany([
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f148"),
        "name" : "Warehouse Jakarta Pusat",
        "type" : "warehouse",
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Pusat",
            "district" : "Gambir",
            "subdistrict" : "Petojo Selatan",
            "zip_code" : "10110",
            "latitude" : -6.1754,
            "longitude" : 106.8272,
            "street_address" : "Jalan Medan Merdeka Barat No.1"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f149"),
        "name" : "Warehouse Jakarta Timur",
        "type" : "warehouse",
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Timur",
            "district" : "Cakung",
            "subdistrict" : "Jatinegara",
            "zip_code" : "13910",
            "latitude" : -6.2211,
            "longitude" : 106.9003,
            "street_address" : "Jalan Raya Bekasi KM 18"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14a"),
        "name" : "Warehouse Tangerang",
        "type" : "warehouse",
        "address" : {
            "province" : "Banten",
            "city" : "Tangerang",
            "district" : "Ciledug",
            "subdistrict" : "Larangan",
            "zip_code" : "15156",
            "latitude" : -6.2294,
            "longitude" : 106.6503,
            "street_address" : "Jalan Haji Agus Salim No.15"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14b"),
        "name" : "Warehouse Serang",
        "type" : "warehouse",
        "address" : {
            "province" : "Banten",
            "city" : "Serang",
            "district" : "Kasemen",
            "subdistrict" : "Banten",
            "zip_code" : "42191",
            "latitude" : -6.1201,
            "longitude" : 106.1502,
            "street_address" : "Jalan Sultan Ageng Tirtayasa No.12"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14c"),
        "name" : "Warehouse Bandung",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bandung",
            "district" : "Cicendo",
            "subdistrict" : "Pasirkaliki",
            "zip_code" : "40171",
            "latitude" : -6.9147,
            "longitude" : 107.6098,
            "street_address" : "Jalan Pasirkaliki No.45"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14d"),
        "name" : "Warehouse Bogor",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bogor",
            "district" : "Bogor Tengah",
            "subdistrict" : "Sempur",
            "zip_code" : "16129",
            "latitude" : -6.595,
            "longitude" : 106.7922,
            "street_address" : "Jalan Pajajaran No.18"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14e"),
        "name" : "Warehouse Semarang",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Semarang",
            "district" : "Semarang Tengah",
            "subdistrict" : "Miroto",
            "zip_code" : "50134",
            "latitude" : -6.9824,
            "longitude" : 110.4091,
            "street_address" : "Jalan Pemuda No.27"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f14f"),
        "name" : "Warehouse Surakarta",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Surakarta",
            "district" : "Laweyan",
            "subdistrict" : "Sriwedari",
            "zip_code" : "57141",
            "latitude" : -7.5589,
            "longitude" : 110.8281,
            "street_address" : "Jalan Slamet Riyadi No.30"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f150"),
        "name" : "Warehouse Yogyakarta",
        "type" : "warehouse",
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Yogyakarta",
            "district" : "Gondokusuman",
            "subdistrict" : "Terban",
            "zip_code" : "55221",
            "latitude" : -7.7828,
            "longitude" : 110.3671,
            "street_address" : "Jalan Jenderal Sudirman No.10"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f151"),
        "name" : "Warehouse Bantul",
        "type" : "warehouse",
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Bantul",
            "district" : "Sewon",
            "subdistrict" : "Timbulharjo",
            "zip_code" : "55751",
            "latitude" : -7.8789,
            "longitude" : 110.3325,
            "street_address" : "Jalan Imogiri Barat KM 6"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f152"),
        "name" : "Warehouse Surabaya",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Surabaya",
            "district" : "Genteng",
            "subdistrict" : "Ketabang",
            "zip_code" : "60272",
            "latitude" : -7.2658,
            "longitude" : 112.743,
            "street_address" : "Jalan Tunjungan No.25"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f153"),
        "name" : "Warehouse Malang",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Malang",
            "district" : "Klojen",
            "subdistrict" : "Kauman",
            "zip_code" : "65111",
            "latitude" : -7.9785,
            "longitude" : 112.6304,
            "street_address" : "Jalan Basuki Rahmat No.16"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f154"),
        "name" : "Warehouse Kediri",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Kediri",
            "district" : "Mojoroto",
            "subdistrict" : "Banjaran",
            "zip_code" : "64111",
            "latitude" : -7.8101,
            "longitude" : 112.0118,
            "street_address" : "Jalan Pahlawan No.20"
        }
    },
    {
        "_id" : ObjectId("67b54e28ba1542ee89a0f155"),
        "name" : "Warehouse Jember",
        "type" : "warehouse",
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Jember",
            "district" : "Patrang",
            "subdistrict" : "Jember Kidul",
            "zip_code" : "68118",
            "latitude" : -8.1688,
            "longitude" : 113.7023,
            "street_address" : "Jalan Ahmad Yani No.33"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f156"),
        "name" : "Depot Jakarta Pusat A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f148"),
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Pusat",
            "district" : "Gambir",
            "subdistrict" : "Petojo Selatan",
            "zip_code" : "10110",
            "latitude" : -6.175,
            "longitude" : 106.8275,
            "street_address" : "Jalan Medan Merdeka Utara No.2"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f157"),
        "name" : "Depot Jakarta Pusat B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f148"),
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Pusat",
            "district" : "Gambir",
            "subdistrict" : "Petojo Selatan",
            "zip_code" : "10110",
            "latitude" : -6.176,
            "longitude" : 106.828,
            "street_address" : "Jalan Veteran No.3"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f158"),
        "name" : "Depot Jakarta Timur A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f149"),
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Timur",
            "district" : "Cakung",
            "subdistrict" : "Jatinegara",
            "zip_code" : "13910",
            "latitude" : -6.222,
            "longitude" : 106.901,
            "street_address" : "Jalan Raya Bekasi KM 19"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f159"),
        "name" : "Depot Jakarta Timur B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f149"),
        "address" : {
            "province" : "DKI Jakarta",
            "city" : "Jakarta Timur",
            "district" : "Cakung",
            "subdistrict" : "Jatinegara",
            "zip_code" : "13910",
            "latitude" : -6.223,
            "longitude" : 106.902,
            "street_address" : "Jalan I Gusti Ngurah Rai No.8"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15a"),
        "name" : "Depot Tangerang A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14a"),
        "address" : {
            "province" : "Banten",
            "city" : "Tangerang",
            "district" : "Ciledug",
            "subdistrict" : "Larangan",
            "zip_code" : "15156",
            "latitude" : -6.23,
            "longitude" : 106.651,
            "street_address" : "Jalan KH. Hasyim Ashari No.12"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15b"),
        "name" : "Depot Tangerang B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14a"),
        "address" : {
            "province" : "Banten",
            "city" : "Tangerang",
            "district" : "Ciledug",
            "subdistrict" : "Larangan",
            "zip_code" : "15156",
            "latitude" : -6.231,
            "longitude" : 106.652,
            "street_address" : "Jalan Gatot Subroto No.20"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15c"),
        "name" : "Depot Serang A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14b"),
        "address" : {
            "province" : "Banten",
            "city" : "Serang",
            "district" : "Kasemen",
            "subdistrict" : "Banten",
            "zip_code" : "42191",
            "latitude" : -6.121,
            "longitude" : 106.151,
            "street_address" : "Jalan Raya Banten Lama No.5"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15d"),
        "name" : "Depot Serang B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14b"),
        "address" : {
            "province" : "Banten",
            "city" : "Serang",
            "district" : "Kasemen",
            "subdistrict" : "Banten",
            "zip_code" : "42191",
            "latitude" : -6.122,
            "longitude" : 106.152,
            "street_address" : "Jalan Sultan Maulana No.3"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15e"),
        "name" : "Depot Bandung A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14c"),
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bandung",
            "district" : "Cicendo",
            "subdistrict" : "Pasirkaliki",
            "zip_code" : "40171",
            "latitude" : -6.915,
            "longitude" : 107.61,
            "street_address" : "Jalan Kebon Kawung No.8"
        }
    },
    {
        "_id" : ObjectId("67b550bfba1542ee89a0f15f"),
        "name" : "Depot Bandung B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14c"),
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bandung",
            "district" : "Cicendo",
            "subdistrict" : "Pasirkaliki",
            "zip_code" : "40171",
            "latitude" : -6.916,
            "longitude" : 107.611,
            "street_address" : "Jalan Pajajaran No.15"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f160"),
        "name" : "Depot Semarang A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14e"),
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Semarang",
            "district" : "Semarang Tengah",
            "subdistrict" : "Miroto",
            "zip_code" : "50134",
            "latitude" : -6.982,
            "longitude" : 110.41,
            "street_address" : "Jalan Pandanaran No.10"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f161"),
        "name" : "Depot Semarang B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14e"),
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Semarang",
            "district" : "Semarang Tengah",
            "subdistrict" : "Miroto",
            "zip_code" : "50134",
            "latitude" : -6.983,
            "longitude" : 110.411,
            "street_address" : "Jalan Gajah Mada No.20"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f162"),
        "name" : "Depot Surakarta A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14f"),
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Surakarta",
            "district" : "Banjarsari",
            "subdistrict" : "Manahan",
            "zip_code" : "57139",
            "latitude" : -7.555,
            "longitude" : 110.831,
            "street_address" : "Jalan Slamet Riyadi No.15"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f163"),
        "name" : "Depot Surakarta B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14f"),
        "address" : {
            "province" : "Jawa Tengah",
            "city" : "Surakarta",
            "district" : "Banjarsari",
            "subdistrict" : "Manahan",
            "zip_code" : "57139",
            "latitude" : -7.556,
            "longitude" : 110.832,
            "street_address" : "Jalan Dr. Radjiman No.25"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f164"),
        "name" : "Depot Yogyakarta A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f150"),
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Yogyakarta",
            "district" : "Gondokusuman",
            "subdistrict" : "Demangan",
            "zip_code" : "55221",
            "latitude" : -7.782,
            "longitude" : 110.383,
            "street_address" : "Jalan Jendral Sudirman No.30"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f165"),
        "name" : "Depot Yogyakarta B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f150"),
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Yogyakarta",
            "district" : "Gondokusuman",
            "subdistrict" : "Demangan",
            "zip_code" : "55221",
            "latitude" : -7.783,
            "longitude" : 110.384,
            "street_address" : "Jalan C. Simanjuntak No.40"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f166"),
        "name" : "Depot Surabaya A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f152"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Surabaya",
            "district" : "Genteng",
            "subdistrict" : "Embong Kaliasin",
            "zip_code" : "60271",
            "latitude" : -7.257,
            "longitude" : 112.752,
            "street_address" : "Jalan Tunjungan No.50"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f167"),
        "name" : "Depot Surabaya B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f152"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Surabaya",
            "district" : "Genteng",
            "subdistrict" : "Embong Kaliasin",
            "zip_code" : "60271",
            "latitude" : -7.258,
            "longitude" : 112.753,
            "street_address" : "Jalan Basuki Rahmat No.60"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f168"),
        "name" : "Depot Malang A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f153"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Malang",
            "district" : "Klojen",
            "subdistrict" : "Kauman",
            "zip_code" : "65119",
            "latitude" : -7.981,
            "longitude" : 112.621,
            "street_address" : "Jalan Kawi No.70"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f169"),
        "name" : "Depot Malang B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f153"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Malang",
            "district" : "Klojen",
            "subdistrict" : "Kauman",
            "zip_code" : "65119",
            "latitude" : -7.982,
            "longitude" : 112.622,
            "street_address" : "Jalan Semeru No.80"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f16a"),
        "name" : "Depot Kediri A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f154"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Kediri",
            "district" : "Pesantren",
            "subdistrict" : "Banjaran",
            "zip_code" : "64131",
            "latitude" : -7.821,
            "longitude" : 112.013,
            "street_address" : "Jalan Hayam Wuruk No.90"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f16b"),
        "name" : "Depot Kediri B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f154"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Kediri",
            "district" : "Pesantren",
            "subdistrict" : "Banjaran",
            "zip_code" : "64131",
            "latitude" : -7.822,
            "longitude" : 112.014,
            "street_address" : "Jalan Brawijaya No.100"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f16c"),
        "name" : "Depot Jember A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f155"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Jember",
            "district" : "Patrang",
            "subdistrict" : "Tegal Besar",
            "zip_code" : "68118",
            "latitude" : -8.168,
            "longitude" : 113.702,
            "street_address" : "Jalan Gajah Mada No.110"
        }
    },
    {
        "_id" : ObjectId("67b550d5ba1542ee89a0f16d"),
        "name" : "Depot Jember B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f155"),
        "address" : {
            "province" : "Jawa Timur",
            "city" : "Jember",
            "district" : "Patrang",
            "subdistrict" : "Tegal Besar",
            "zip_code" : "68118",
            "latitude" : -8.169,
            "longitude" : 113.703,
            "street_address" : "Jalan Sultan Agung No.120"
        }
    },
    {
        "_id" : ObjectId("67b55563ba1542ee89a0f16e"),
        "name" : "Depot Bantul A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f151"),
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Bantul",
            "district" : "Bantul",
            "subdistrict" : "Trirenggo",
            "zip_code" : "55714",
            "latitude" : -7.8889,
            "longitude" : 110.328,
            "street_address" : "Jl. Bantul No. 20"
        }
    },
    {
        "_id" : ObjectId("67b55563ba1542ee89a0f16f"),
        "name" : "Depot Bantul B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f151"),
        "address" : {
            "province" : "DI Yogyakarta",
            "city" : "Bantul",
            "district" : "Sewon",
            "subdistrict" : "Pendowoharjo",
            "zip_code" : "55185",
            "latitude" : -7.8254,
            "longitude" : 110.3492,
            "street_address" : "Jl. Parangtritis No. 30"
        }
    },
    {
        "_id" : ObjectId("67b55563ba1542ee89a0f170"),
        "name" : "Depot Bogor A",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14d"),
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bogor",
            "district" : "Bogor Selatan",
            "subdistrict" : "Cikaret",
            "zip_code" : "16132",
            "latitude" : -6.6145,
            "longitude" : 106.8063,
            "street_address" : "Jl. Raya Cikaret No. 15"
        }
    },
    {
        "_id" : ObjectId("67b55563ba1542ee89a0f171"),
        "name" : "Depot Bogor B",
        "type" : "depot",
        "warehouse_id" : ObjectId("67b54e28ba1542ee89a0f14d"),
        "address" : {
            "province" : "Jawa Barat",
            "city" : "Bogor",
            "district" : "Bogor Tengah",
            "subdistrict" : "Pabaton",
            "zip_code" : "16121",
            "latitude" : -6.5902,
            "longitude" : 106.7995,
            "street_address" : "Jl. Juanda No. 45"
        }
    }
]);

print("Database seeded successfully!");
EOF

if [ $? -eq 0 ]; then
    echo "Database seeded successfully!"
else
    echo "Error during database seeding!"
    exit 1
fi
