package wrapper

// #cgo LDFLAGS: -lrsync
// #include <librsync.h>
import "C"
import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// /** Generate the signature of a basis file, and write it out to another.
// *
// * It's recommended you use rs_sig_args() to get the recommended arguments for
// * this based on the original file size.
// *
// * \param old_file Stdio readable file whose signature will be generated.
// *
// * \param sig_file Writable stdio file to which the signature will be written./
// *
// * \param block_len Checksum block size to use (0 for "recommended"). Larger
// * values make the signature shorter, and the delta longer.
// *
// * \param strong_len Strongsum length in bytes to use (0 for "maximum", -1 for
// * "minimum"). Smaller values make the signature shorter but increase the risk
// * of corruption from hash collisions.
// *
// * \param sig_magic Signature file format to generate (0 for "recommended").
// * See ::rs_magic_number.
// *
// * \param stats Optional pointer to receive statistics.
// *
// * \sa \ref api_whole */
// rs_result rs_sig_file(FILE *old_file, FILE *sig_file,
//                                      size_t block_len, size_t strong_len,
//                                      rs_magic_number sig_magic,
//                                      rs_stats_t *stats);

func RSSig(file *os.File) *os.File {
	sig, err := ioutil.TempFile("", filepath.Base(file.Name()))
	if err != nil {
		return nil
	}

	oldFile := C.fdopen(C.int(file.Fd()), C.CString("r"))
	sigFile := C.fdopen(C.int(sig.Fd()), C.CString("w+"))
	var stats C.rs_stats_t
	ret := C.rs_sig_file(oldFile, sigFile, 0, 0, 0, &stats)
	if ret == C.RS_DONE {
		return sig
	}

	sig.Close()
	return nil
}

//
///** Load signatures from a signature file into memory.
// *
// * \param sig_file Readable stdio file from which the signature will be read.
// *
// * \param sumset on return points to the newly allocated structure.
// *
// * \param stats Optional pointer to receive statistics.
// *
// * \sa \ref api_whole */
// rs_result rs_loadsig_file(FILE *sig_file,
//                                          rs_signature_t **sumset,
//                                          rs_stats_t *stats);
//
///** Generate a delta between a signature and a new file into a delta file.
// *
// * \sa \ref api_whole */
// rs_result rs_delta_file(rs_signature_t *, FILE *new_file,
//                                        FILE *delta_file, rs_stats_t *);
//
///** Apply a patch, relative to a basis, into a new file.
// *
// * \sa \ref api_whole */
// rs_result rs_patch_file(FILE *basis_file, FILE *delta_file,
//                                        FILE *new_file, rs_stats_t *);
