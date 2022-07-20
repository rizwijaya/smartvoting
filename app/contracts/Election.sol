// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.9.0;
//pragma abicoder v2;

contract Election {
    address public admin;
    uint256 candidateCount;
    uint256 voterCount;
    bool start;
    bool end;
    // Initilizing default values
    constructor() { 
        admin = msg.sender;
        candidateCount = 0;
        voterCount = 0;
        start = false;
        end = false;
    }
    //------Model data------//
    // Model data Kandidat
    struct Candidate {
        uint256 candidateId;
        string nama;
        string deskripsi;
        string slogan;
        uint256 voteCount;
    }
    // Model data Detail Pemilihan
    struct ElectionDetails {
        string adminNama;
        string adminEmail;
        string adminTitle;
        string electionDeskripsi;
        string electionTitle;
        string organizationTitle;
    }
    // Modeling data Pemilih
    struct Voter {
        address voterAddress;
        string name;
        string home_address;
        string phone;
        bool isVerified;
        bool hasVoted;
        bool isRegistered;
    }
    //------end Model data------//
   
    // Permission hanya admin yang dapat mengakses
    modifier onlyAdmin() {
        require(msg.sender == admin);
        _;
    }

    //------Candidate------//
    mapping(uint256 => Candidate) public detailKandidat;
    // Tambahkan Candidate baru
    function post_AdminTambahKandidat(string memory _nama, string memory _deskripsi, string memory _slogan)
        public
        // Only admin can add
        onlyAdmin
    {
        Candidate memory newCandidate =
            Candidate({
                candidateId: candidateCount,
                deskripsi: _deskripsi,
                nama: _nama,
                slogan: _slogan,
                voteCount: 0
            });
        detailKandidat[candidateCount] = newCandidate;
        candidateCount += 1;
    }
    //------end Candidate------//

    //------Admin------//
    ElectionDetails electionDetails;

    function post_AdminCreatePemilihan(
        string memory _adminNama,
        string memory _adminEmail,
        string memory _adminTitle,
        string memory _electionDeskripsi,
        string memory _electionTitle,
        string memory _organizationTitle
    )
        public
        // Only admin can add
        onlyAdmin
    {
        electionDetails = ElectionDetails(
            _adminNama,
            _adminEmail,
            _adminTitle,
            _electionDeskripsi,
            _electionTitle,
            _organizationTitle
        );
    }
    
    // Daftarkan Pemilih baru dari Admin
    function post_AdminPemilihBaru(address _address_public, string memory _name, string memory _home_address, string memory _phone) public {
        Voter memory newVoter =
            Voter({
                voterAddress: _address_public,
                name: _name,
                home_address: _home_address,
                phone: _phone,
                hasVoted: false,
                isVerified: false,
                isRegistered: true
            });
        detailPemilih[msg.sender] = newVoter;
        voters.push(msg.sender);
        voterCount += 1;
    }

    //------Detail Admin------//
    // function get_DetailAdmin() public view returns (ElectionDetails memory) {
    //     return electionDetails;
    // }
    // Dapatkan Address dari admin / user
    function get_Admin() public view returns (address) {
        return admin;
    }
    // Dapatkan nama dari admin / user
    function get_AdminName() public view returns (string memory) {
        return electionDetails.adminNama;
    }
    // Dapatkan email dari admin / user
    function get_AdminEmail() public view returns (string memory) {
        return electionDetails.adminEmail;
    }
    //Dapatkan title dari admin / user
    function get_AdminTitle() public view returns (string memory) {
        return electionDetails.adminTitle;
    }
    //------Detail Pemilihan------//
    //Dapatkan title pemilihan
    function get_PemilihanTitle() public view returns (string memory) {
        return electionDetails.electionTitle;
    }
    //Dapatkan deskripsi pemilihan
    function get_PemilihanDeskripsi() public view returns (string memory) {
        return electionDetails.electionDeskripsi;
    }
    //Dapatkan Organisasi Penyelanggara Pemilihan
    function get_Organisasi() public view returns (string memory) {
        return electionDetails.organizationTitle;
    }
    //------end Admin------//

    //------Pemilih------//
    address[] public voters;
    mapping(address => Voter) public detailPemilih;

    // Daftarkan Pemilih Baru
    function post_RegisterPemilih(string memory _name, string memory _home_address, string memory _phone) public {
        Voter memory newVoter =
            Voter({
                voterAddress: msg.sender,
                name: _name,
                home_address: _home_address,
                phone: _phone,
                hasVoted: false,
                isVerified: false,
                isRegistered: true
            });
        detailPemilih[msg.sender] = newVoter;
        voters.push(msg.sender);
        voterCount += 1;
    }

    // Verifikasi Pemilih
    function post_AdminVerifikasiPemilih(bool _verifedStatus, address voterAddress)
        public
        // Only admin can verify
        onlyAdmin
    {
        detailPemilih[voterAddress].isVerified = _verifedStatus;
    }
    //------end Pemilih------//

    //------Proses Pemilihan------//
    // Dapatkan Jumlah Candidate
    function get_TotalKandidat() public view returns (uint256) {
        return candidateCount;
    }
    // Dapatkan Jumlah pemilih
    function get_TotalPemilih() public view returns (uint256) {
        return voterCount;
    }
    
    //Lakukan Pemilihan
    function post_GeneralVote(uint256 candidateId) public {
        require(detailPemilih[msg.sender].hasVoted == false);
        require(detailPemilih[msg.sender].isVerified == true);
        require(start == true);
        require(end == false);
        detailKandidat[candidateId].voteCount += 1;
        detailPemilih[msg.sender].hasVoted = true;
    }

    //Mulai Pemilihan
    function mulaiPemilihan() public onlyAdmin {
        start = true;
        end = false;
    }
    // Akhiri Pemilihan
    function akhiriPemilihan() public onlyAdmin {
        end = true;
        start = false;
    }

    // Dapatkan status Mulai Pemilihan
    function get_MulaiPemilihan() public view returns (bool) {
        return start;
    }

    // Dapatkan status Pemilihan Berakhir
    function get_AkhiriPemilihan() public view returns (bool) {
        return end;
    }
    //------end Proses Pemilihan------//
}
