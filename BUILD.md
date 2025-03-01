# Build for 1.24-master-039b3ebeba
Quick 2 step build guide with output.
```bash
@melinko2003 ➜ /workspaces/shim-go/src/go (master) $ cd src/
# Build with Full testing and Public API verification
@melinko2003 ➜ .../shim-go/src/go/src (master) $ ./all.bash 
# Build Output
Building Go cmd/dist using /usr/local/go. (go1.23.4 linux/amd64)
Building Go toolchain1 using /usr/local/go.
Building Go bootstrap cmd/go (go_bootstrap) using Go toolchain1.
Building Go toolchain2 using go_bootstrap and Go toolchain1.
Building Go toolchain3 using go_bootstrap and Go toolchain2.
Building packages and commands for linux/amd64.

##### Test execution environment.
# GOARCH: amd64
# CPU: AMD EPYC 7763 64-Core Processor                
# GOOS: linux
# OS Version: Linux 6.5.0-1025-azure #26~22.04.1-Ubuntu SMP Thu Jul 11 22:33:04 UTC 2024 x86_64

##### Testing packages.
ok      archive/tar     0.195s
ok      archive/zip     0.146s
ok      bufio   0.126s
ok      bytes   0.411s
ok      cmp     0.027s
ok      compress/bzip2  0.056s
ok      compress/flate  0.744s
ok      compress/gzip   1.733s
ok      compress/lzw    0.009s
ok      compress/zlib   0.013s
ok      container/heap  0.003s
ok      container/list  0.002s
ok      container/ring  0.013s
ok      context 0.023s
ok      crypto  0.767s
ok      crypto/aes      0.006s
ok      crypto/cipher   4.132s
ok      crypto/des      0.007s
ok      crypto/dsa      0.003s
ok      crypto/ecdh     0.044s
ok      crypto/ecdsa    0.042s
ok      crypto/ed25519  0.049s
ok      crypto/elliptic 0.037s
?       crypto/fips140  [no test files]
ok      crypto/hkdf     0.005s
ok      crypto/hmac     0.011s
ok      crypto/internal/boring  0.002s
?       crypto/internal/boring/bbig     [no test files]
ok      crypto/internal/boring/bcache   0.150s
?       crypto/internal/boring/sig      [no test files]
?       crypto/internal/cryptotest      [no test files]
?       crypto/internal/entropy [no test files]
?       crypto/internal/fips140 [no test files]
ok      crypto/internal/fips140/aes     0.018s
ok      crypto/internal/fips140/aes/gcm 0.007s [no tests to run]
?       crypto/internal/fips140/alias   [no test files]
ok      crypto/internal/fips140/bigmod  0.067s
?       crypto/internal/fips140/check   [no test files]
?       crypto/internal/fips140/check/checktest [no test files]
ok      crypto/internal/fips140/drbg    0.002s [no tests to run]
ok      crypto/internal/fips140/ecdh    0.002s
ok      crypto/internal/fips140/ecdsa   0.058s
?       crypto/internal/fips140/ed25519 [no test files]
ok      crypto/internal/fips140/edwards25519    0.075s
ok      crypto/internal/fips140/edwards25519/field      0.013s
?       crypto/internal/fips140/hkdf    [no test files]
?       crypto/internal/fips140/hmac    [no test files]
ok      crypto/internal/fips140/mlkem   0.270s
ok      crypto/internal/fips140/nistec  0.002s
ok      crypto/internal/fips140/nistec/fiat     0.002s [no tests to run]
?       crypto/internal/fips140/pbkdf2  [no test files]
ok      crypto/internal/fips140/rsa     0.081s
?       crypto/internal/fips140/sha256  [no test files]
?       crypto/internal/fips140/sha3    [no test files]
?       crypto/internal/fips140/sha512  [no test files]
?       crypto/internal/fips140/ssh     [no test files]
ok      crypto/internal/fips140/subtle  0.002s
?       crypto/internal/fips140/tls12   [no test files]
?       crypto/internal/fips140/tls13   [no test files]
ok      crypto/internal/fips140deps     0.082s
?       crypto/internal/fips140deps/byteorder   [no test files]
?       crypto/internal/fips140deps/cpu [no test files]
?       crypto/internal/fips140deps/godebug     [no test files]
?       crypto/internal/fips140hash     [no test files]
?       crypto/internal/fips140only     [no test files]
ok      crypto/internal/fips140test     0.853s
ok      crypto/internal/hpke    0.004s
?       crypto/internal/impl    [no test files]
?       crypto/internal/randutil        [no test files]
ok      crypto/internal/sysrand 0.263s
?       crypto/internal/sysrand/internal/seccomp        [no test files]
ok      crypto/md5      0.005s
ok      crypto/mlkem    0.069s
ok      crypto/pbkdf2   0.018s
ok      crypto/rand     0.450s
ok      crypto/rc4      0.030s
ok      crypto/rsa      0.445s
ok      crypto/sha1     0.021s
ok      crypto/sha256   0.005s
ok      crypto/sha3     1.993s
ok      crypto/sha512   0.007s
ok      crypto/subtle   0.218s
ok      crypto/tls      7.732s
?       crypto/tls/internal/fips140tls  [no test files]
ok      crypto/x509     0.493s
?       crypto/x509/pkix        [no test files]
ok      database/sql    0.497s
ok      database/sql/driver     0.002s
ok      debug/buildinfo 0.045s
ok      debug/dwarf     0.012s
ok      debug/elf       0.178s
ok      debug/gosym     0.222s
ok      debug/macho     0.005s
ok      debug/pe        0.036s
ok      debug/plan9obj  0.002s
ok      embed   0.007s [no tests to run]
ok      embed/internal/embedtest        0.005s
?       encoding        [no test files]
ok      encoding/ascii85        0.004s
ok      encoding/asn1   0.003s
ok      encoding/base32 0.036s
ok      encoding/base64 0.005s
ok      encoding/binary 0.005s
ok      encoding/csv    0.012s
ok      encoding/gob    3.619s
ok      encoding/hex    0.003s
ok      encoding/json   0.228s
ok      encoding/pem    0.566s
ok      encoding/xml    0.040s
ok      errors  0.018s
ok      expvar  0.021s
ok      flag    0.023s
ok      fmt     0.058s
ok      go/ast  0.007s
ok      go/ast/internal/tests   0.009s
ok      go/build        4.717s
ok      go/build/constraint     0.067s
ok      go/constant     0.003s
ok      go/doc  0.038s
ok      go/doc/comment  0.766s
ok      go/format       0.004s
ok      go/importer     0.233s
ok      go/internal/gccgoimporter       0.009s
ok      go/internal/gcimporter  2.060s
ok      go/internal/srcimporter 9.412s
ok      go/parser       0.155s
ok      go/printer      0.165s
ok      go/scanner      0.003s
ok      go/token        0.013s
ok      go/types        6.509s
ok      go/version      0.002s
ok      hash    0.002s
ok      hash/adler32    0.008s
ok      hash/crc32      0.006s
ok      hash/crc64      0.002s
ok      hash/fnv        0.003s
ok      hash/maphash    0.146s
ok      html    0.010s
ok      html/template   0.063s
ok      image   0.051s
ok      image/color     0.018s
?       image/color/palette     [no test files]
ok      image/draw      0.081s
ok      image/gif       0.203s
?       image/internal/imageutil        [no test files]
ok      image/jpeg      0.172s
ok      image/png       0.333s
ok      index/suffixarray       0.158s
ok      internal/abi    0.047s
?       internal/asan   [no test files]
?       internal/bisect [no test files]
ok      internal/buildcfg       0.002s
?       internal/bytealg        [no test files]
?       internal/byteorder      [no test files]
?       internal/cfg    [no test files]
ok      internal/chacha8rand    0.002s
ok      internal/copyright      0.317s
?       internal/coverage       [no test files]
?       internal/coverage/calloc        [no test files]
ok      internal/coverage/cfile 0.788s
ok      internal/coverage/cformat       0.011s
ok      internal/coverage/cmerge        0.008s
?       internal/coverage/decodecounter [no test files]
?       internal/coverage/decodemeta    [no test files]
?       internal/coverage/encodecounter [no test files]
?       internal/coverage/encodemeta    [no test files]
ok      internal/coverage/pods  0.004s
?       internal/coverage/rtcov [no test files]
ok      internal/coverage/slicereader   0.002s
ok      internal/coverage/slicewriter   0.002s
?       internal/coverage/stringtab     [no test files]
ok      internal/coverage/test  0.007s
?       internal/coverage/uleb128       [no test files]
ok      internal/cpu    0.014s
ok      internal/dag    0.015s
ok      internal/diff   0.003s
?       internal/exportdata     [no test files]
?       internal/filepathlite   [no test files]
ok      internal/fmtsort        0.002s
ok      internal/fuzz   0.006s
?       internal/goarch [no test files]
ok      internal/godebug        0.394s
ok      internal/godebugs       9.019s
?       internal/goexperiment   [no test files]
?       internal/goos   [no test files]
?       internal/goroot [no test files]
ok      internal/gover  0.005s
?       internal/goversion      [no test files]
ok      internal/itoa   0.002s
?       internal/lazyregexp     [no test files]
?       internal/lazytemplate   [no test files]
?       internal/msan   [no test files]
?       internal/nettrace       [no test files]
?       internal/obscuretestdata        [no test files]
?       internal/oserror        [no test files]
ok      internal/pkgbits        0.002s
ok      internal/platform       0.999s
ok      internal/poll   0.137s
ok      internal/profile        0.003s
?       internal/profilerecord  [no test files]
?       internal/race   [no test files]
ok      internal/reflectlite    0.003s
ok      internal/runtime/atomic 0.013s
?       internal/runtime/exithook       [no test files]
ok      internal/runtime/maps   0.004s
ok      internal/runtime/math   0.002s
ok      internal/runtime/sys    0.002s
ok      internal/runtime/syscall        0.002s
ok      internal/saferio        0.017s
ok      internal/singleflight   0.014s
?       internal/stringslite    [no test files]
ok      internal/sync   0.583s
ok      internal/synctest       0.900s
?       internal/syscall/execenv        [no test files]
ok      internal/syscall/unix   0.009s
ok      internal/sysinfo        0.002s
?       internal/syslist        [no test files]
ok      internal/testenv        0.240s
?       internal/testlog        [no test files]
?       internal/testpty        [no test files]
ok      internal/trace  27.591s
?       internal/trace/internal/testgen [no test files]
ok      internal/trace/internal/tracev1 0.260s
?       internal/trace/raw      [no test files]
?       internal/trace/testtrace        [no test files]
ok      internal/trace/tracev2  0.009s
?       internal/trace/traceviewer      [no test files]
?       internal/trace/traceviewer/format       [no test files]
?       internal/trace/version  [no test files]
?       internal/txtar  [no test files]
ok      internal/types/errors   0.466s
ok      internal/unsafeheader   0.012s
ok      internal/xcoff  0.017s
ok      internal/zstd   0.939s
ok      io      0.024s
ok      io/fs   0.384s
ok      io/ioutil       0.008s
ok      iter    0.013s
ok      log     0.004s
?       log/internal    [no test files]
ok      log/slog        0.028s
?       log/slog/internal       [no test files]
ok      log/slog/internal/benchmarks    0.002s
ok      log/slog/internal/buffer        0.002s
?       log/slog/internal/slogtest      [no test files]
ok      log/syslog      1.214s
ok      maps    0.002s
ok      math    0.005s
ok      math/big        2.227s
ok      math/bits       0.009s
ok      math/cmplx      0.012s
ok      math/rand       0.084s
ok      math/rand/v2    0.294s
ok      mime    0.016s
ok      mime/multipart  1.023s
ok      mime/quotedprintable    0.034s
ok      net     6.226s
ok      net/http        11.660s
ok      net/http/cgi    0.763s
ok      net/http/cookiejar      0.020s
ok      net/http/fcgi   0.217s
ok      net/http/httptest       0.029s
ok      net/http/httptrace      0.004s
ok      net/http/httputil       0.635s
ok      net/http/internal       0.092s
ok      net/http/internal/ascii 0.002s
?       net/http/internal/httpcommon    [no test files]
?       net/http/internal/testcert      [no test files]
ok      net/http/pprof  5.098s
ok      net/internal/cgotest    0.001s
ok      net/internal/socktest   0.002s
ok      net/mail        0.004s
ok      net/netip       0.446s
ok      net/rpc 0.027s
ok      net/rpc/jsonrpc 0.008s
ok      net/smtp        0.012s
ok      net/textproto   0.015s
ok      net/url 0.007s
ok      os      1.490s
ok      os/exec 0.487s
ok      os/exec/internal/fdtest 0.002s
ok      os/signal       3.020s
ok      os/user 0.003s
ok      path    0.003s
ok      path/filepath   0.032s
ok      plugin  0.012s
ok      reflect 0.614s
?       reflect/internal/example1       [no test files]
?       reflect/internal/example2       [no test files]
ok      regexp  0.500s
ok      regexp/syntax   0.943s
ok      runtime 76.195s
ok      runtime/cgo     0.003s
?       runtime/coverage        [no test files]
ok      runtime/debug   0.060s
?       runtime/internal/startlinetest  [no test files]
ok      runtime/internal/wasitest       0.016s
ok      runtime/metrics 0.010s
ok      runtime/pprof   20.229s
?       runtime/race    [no test files]
?       runtime/race/internal/amd64v1   [no test files]
ok      runtime/trace   0.104s
ok      slices  0.052s
ok      sort    0.070s
ok      strconv 0.298s
ok      strings 0.147s
?       structs [no test files]
ok      sync    0.339s
ok      sync/atomic     0.731s
ok      syscall 3.659s
ok      testing 1.255s
ok      testing/fstest  0.004s
?       testing/internal/testdeps       [no test files]
ok      testing/iotest  0.002s
ok      testing/quick   0.030s
ok      testing/slogtest        0.013s
ok      text/scanner    0.003s
ok      text/tabwriter  0.003s
ok      text/template   0.027s
ok      text/template/parse     0.005s
ok      time    10.786s
?       time/tzdata     [no test files]
ok      unicode 0.003s
ok      unicode/utf16   0.002s
ok      unicode/utf8    0.004s
ok      unique  0.075s
?       unsafe  [no test files]
ok      weak    0.022s
ok      cmd/addr2line   0.675s
ok      cmd/api 17.609s
?       cmd/asm [no test files]
?       cmd/asm/internal/arch   [no test files]
ok      cmd/asm/internal/asm    0.553s
?       cmd/asm/internal/flags  [no test files]
ok      cmd/asm/internal/lex    0.002s
?       cmd/buildid     [no test files]
?       cmd/cgo [no test files]
?       cmd/cgo/internal/cgotest        [no test files]
ok      cmd/cgo/internal/swig   0.009s
ok      cmd/cgo/internal/test   1.068s
?       cmd/cgo/internal/test/gcc68255  [no test files]
?       cmd/cgo/internal/test/issue23555a       [no test files]
?       cmd/cgo/internal/test/issue23555b       [no test files]
?       cmd/cgo/internal/test/issue26213        [no test files]
?       cmd/cgo/internal/test/issue26430        [no test files]
?       cmd/cgo/internal/test/issue26743        [no test files]
?       cmd/cgo/internal/test/issue27054        [no test files]
?       cmd/cgo/internal/test/issue27340        [no test files]
?       cmd/cgo/internal/test/issue29563        [no test files]
?       cmd/cgo/internal/test/issue30527        [no test files]
?       cmd/cgo/internal/test/issue41761a       [no test files]
?       cmd/cgo/internal/test/issue43639        [no test files]
?       cmd/cgo/internal/test/issue52611a       [no test files]
?       cmd/cgo/internal/test/issue52611b       [no test files]
?       cmd/cgo/internal/test/issue8756 [no test files]
?       cmd/cgo/internal/test/issue8828 [no test files]
?       cmd/cgo/internal/test/issue9026 [no test files]
?       cmd/cgo/internal/test/issue9400 [no test files]
?       cmd/cgo/internal/test/issue9510a        [no test files]
?       cmd/cgo/internal/test/issue9510b        [no test files]
ok      cmd/cgo/internal/testcarchive   0.003s
ok      cmd/cgo/internal/testcshared    0.002s
ok      cmd/cgo/internal/testerrors     45.189s
ok      cmd/cgo/internal/testfortran    0.009s
ok      cmd/cgo/internal/testgodefs     0.748s
ok      cmd/cgo/internal/testlife       1.231s
ok      cmd/cgo/internal/testnocgo      0.004s
ok      cmd/cgo/internal/testplugin     0.005s
ok      cmd/cgo/internal/testsanitizers 86.024s
ok      cmd/cgo/internal/testshared     0.021s
ok      cmd/cgo/internal/testso 5.156s
ok      cmd/cgo/internal/teststdio      4.160s
ok      cmd/cgo/internal/testtls        0.029s
ok      cmd/compile     12.196s
?       cmd/compile/internal/abi        [no test files]
ok      cmd/compile/internal/abt        0.034s
ok      cmd/compile/internal/amd64      5.970s
?       cmd/compile/internal/arm        [no test files]
?       cmd/compile/internal/arm64      [no test files]
ok      cmd/compile/internal/base       0.002s
?       cmd/compile/internal/bitvec     [no test files]
ok      cmd/compile/internal/compare    0.003s
?       cmd/compile/internal/coverage   [no test files]
?       cmd/compile/internal/deadlocals [no test files]
ok      cmd/compile/internal/devirtualize       0.006s
ok      cmd/compile/internal/dwarfgen   0.545s
?       cmd/compile/internal/escape     [no test files]
?       cmd/compile/internal/gc [no test files]
ok      cmd/compile/internal/importer   1.554s
?       cmd/compile/internal/inline     [no test files]
ok      cmd/compile/internal/inline/inlheur     0.469s
?       cmd/compile/internal/inline/interleaved [no test files]
ok      cmd/compile/internal/ir 0.002s
ok      cmd/compile/internal/liveness   0.022s
ok      cmd/compile/internal/logopt     0.245s
?       cmd/compile/internal/loong64    [no test files]
ok      cmd/compile/internal/loopvar    55.565s
?       cmd/compile/internal/mips       [no test files]
?       cmd/compile/internal/mips64     [no test files]
ok      cmd/compile/internal/noder      0.004s
?       cmd/compile/internal/objw       [no test files]
?       cmd/compile/internal/pgoir      [no test files]
?       cmd/compile/internal/pkginit    [no test files]
?       cmd/compile/internal/ppc64      [no test files]
ok      cmd/compile/internal/rangefunc  0.005s
ok      cmd/compile/internal/reflectdata        0.005s [no tests to run]
?       cmd/compile/internal/riscv64    [no test files]
?       cmd/compile/internal/rttype     [no test files]
?       cmd/compile/internal/s390x      [no test files]
ok      cmd/compile/internal/ssa        59.525s
ok      cmd/compile/internal/ssagen     0.016s
?       cmd/compile/internal/staticdata [no test files]
?       cmd/compile/internal/staticinit [no test files]
ok      cmd/compile/internal/syntax     0.267s
ok      cmd/compile/internal/test       13.949s
?       cmd/compile/internal/typebits   [no test files]
ok      cmd/compile/internal/typecheck  0.564s
ok      cmd/compile/internal/types      0.002s
ok      cmd/compile/internal/types2     6.380s
?       cmd/compile/internal/walk       [no test files]
?       cmd/compile/internal/wasm       [no test files]
?       cmd/compile/internal/x86        [no test files]
ok      cmd/covdata     0.007s
ok      cmd/cover       2.046s
ok      cmd/dist        0.019s
ok      cmd/distpack    0.015s
ok      cmd/doc 0.185s
ok      cmd/fix 6.753s
ok      cmd/go  149.157s
ok      cmd/go/internal/auth    0.039s
?       cmd/go/internal/base    [no test files]
?       cmd/go/internal/bug     [no test files]
ok      cmd/go/internal/cache   0.549s
?       cmd/go/internal/cacheprog       [no test files]
ok      cmd/go/internal/cfg     0.009s [no tests to run]
?       cmd/go/internal/clean   [no test files]
?       cmd/go/internal/cmdflag [no test files]
?       cmd/go/internal/doc     [no test files]
ok      cmd/go/internal/envcmd  0.026s
ok      cmd/go/internal/fips140 0.017s
?       cmd/go/internal/fix     [no test files]
?       cmd/go/internal/fmtcmd  [no test files]
ok      cmd/go/internal/fsys    0.063s
ok      cmd/go/internal/generate        0.011s
ok      cmd/go/internal/gover   0.012s
?       cmd/go/internal/help    [no test files]
ok      cmd/go/internal/imports 0.013s
?       cmd/go/internal/list    [no test files]
ok      cmd/go/internal/load    0.014s
ok      cmd/go/internal/lockedfile      0.283s
ok      cmd/go/internal/lockedfile/internal/filelock    0.063s
ok      cmd/go/internal/mmap    0.004s
?       cmd/go/internal/modcmd  [no test files]
ok      cmd/go/internal/modfetch        0.021s
ok      cmd/go/internal/modfetch/codehost       2.835s
ok      cmd/go/internal/modfetch/zip_sum_test   0.040s
?       cmd/go/internal/modget  [no test files]
ok      cmd/go/internal/modindex        0.251s
?       cmd/go/internal/modinfo [no test files]
ok      cmd/go/internal/modload 0.010s
ok      cmd/go/internal/mvs     0.022s
?       cmd/go/internal/run     [no test files]
?       cmd/go/internal/search  [no test files]
ok      cmd/go/internal/str     0.003s
?       cmd/go/internal/telemetrycmd    [no test files]
?       cmd/go/internal/telemetrystats  [no test files]
ok      cmd/go/internal/test    0.094s
?       cmd/go/internal/test/internal/genflags  [no test files]
?       cmd/go/internal/tool    [no test files]
ok      cmd/go/internal/toolchain       0.016s
?       cmd/go/internal/trace   [no test files]
ok      cmd/go/internal/vcs     0.025s
ok      cmd/go/internal/vcweb   0.017s
ok      cmd/go/internal/vcweb/vcstest   3.014s
?       cmd/go/internal/version [no test files]
?       cmd/go/internal/vet     [no test files]
ok      cmd/go/internal/web     0.005s
?       cmd/go/internal/web/intercept   [no test files]
ok      cmd/go/internal/work    0.233s
?       cmd/go/internal/workcmd [no test files]
ok      cmd/gofmt       0.049s
ok      cmd/internal/archive    1.075s
?       cmd/internal/bio        [no test files]
ok      cmd/internal/bootstrap_test     0.012s
?       cmd/internal/browser    [no test files]
ok      cmd/internal/buildid    0.239s
?       cmd/internal/codesign   [no test files]
ok      cmd/internal/cov        0.522s
?       cmd/internal/cov/covcmd [no test files]
?       cmd/internal/disasm     [no test files]
ok      cmd/internal/dwarf      0.016s
ok      cmd/internal/edit       0.007s
?       cmd/internal/gcprog     [no test files]
ok      cmd/internal/goobj      0.014s
?       cmd/internal/hash       [no test files]
?       cmd/internal/macho      [no test files]
ok      cmd/internal/moddeps    2.348s
ok      cmd/internal/obj        0.962s
?       cmd/internal/obj/arm    [no test files]
ok      cmd/internal/obj/arm64  0.511s
ok      cmd/internal/obj/loong64        0.080s
?       cmd/internal/obj/mips   [no test files]
ok      cmd/internal/obj/ppc64  0.176s
ok      cmd/internal/obj/riscv  0.083s
ok      cmd/internal/obj/s390x  0.005s
?       cmd/internal/obj/wasm   [no test files]
ok      cmd/internal/obj/x86    0.303s
ok      cmd/internal/objabi     0.019s
?       cmd/internal/objfile    [no test files]
ok      cmd/internal/osinfo     0.003s
ok      cmd/internal/par        0.007s
?       cmd/internal/pathcache  [no test files]
ok      cmd/internal/pgo        0.014s
ok      cmd/internal/pkgpath    0.031s
ok      cmd/internal/pkgpattern 0.006s
ok      cmd/internal/quoted     0.013s
?       cmd/internal/robustio   [no test files]
?       cmd/internal/script     [no test files]
?       cmd/internal/script/scripttest  [no test files]
ok      cmd/internal/src        0.003s
ok      cmd/internal/sys        0.005s
?       cmd/internal/telemetry  [no test files]
?       cmd/internal/telemetry/counter  [no test files]
ok      cmd/internal/test2json  0.177s
ok      cmd/link        61.572s
?       cmd/link/internal/amd64 [no test files]
?       cmd/link/internal/arm   [no test files]
?       cmd/link/internal/arm64 [no test files]
ok      cmd/link/internal/benchmark     0.027s
?       cmd/link/internal/dwtest        [no test files]
ok      cmd/link/internal/ld    21.080s
?       cmd/link/internal/loadelf       [no test files]
ok      cmd/link/internal/loader        0.043s
?       cmd/link/internal/loadmacho     [no test files]
?       cmd/link/internal/loadpe        [no test files]
?       cmd/link/internal/loadxcoff     [no test files]
?       cmd/link/internal/loong64       [no test files]
?       cmd/link/internal/mips  [no test files]
?       cmd/link/internal/mips64        [no test files]
?       cmd/link/internal/ppc64 [no test files]
?       cmd/link/internal/riscv64       [no test files]
?       cmd/link/internal/s390x [no test files]
?       cmd/link/internal/sym   [no test files]
?       cmd/link/internal/wasm  [no test files]
?       cmd/link/internal/x86   [no test files]
ok      cmd/nm  3.121s
ok      cmd/objdump     5.003s
ok      cmd/pack        0.522s
ok      cmd/pprof       1.492s
?       cmd/preprofile  [no test files]
ok      cmd/relnote     0.002s
?       cmd/test2json   [no test files]
ok      cmd/trace       0.063s
ok      cmd/vet 9.393s

##### os/user with tag osusergo
ok      os/user 0.031s

##### hash/maphash purego implementation
ok      hash/maphash    3.524s

##### crypto with tag purego (build and vet only)
?       crypto/fips140  [no test files]
?       crypto/internal/boring/bbig     [no test files]
?       crypto/internal/boring/sig      [no test files]
?       crypto/internal/cryptotest      [no test files]
?       crypto/internal/entropy [no test files]
?       crypto/internal/fips140 [no test files]
?       crypto/internal/fips140/alias   [no test files]
?       crypto/internal/fips140/check   [no test files]
?       crypto/internal/fips140/check/checktest [no test files]
?       crypto/internal/fips140/ed25519 [no test files]
?       crypto/internal/fips140/hkdf    [no test files]
?       crypto/internal/fips140/hmac    [no test files]
?       crypto/internal/fips140/pbkdf2  [no test files]
?       crypto/internal/fips140/sha256  [no test files]
?       crypto/internal/fips140/sha3    [no test files]
?       crypto/internal/fips140/sha512  [no test files]
?       crypto/internal/fips140/ssh     [no test files]
?       crypto/internal/fips140/tls12   [no test files]
?       crypto/internal/fips140/tls13   [no test files]
?       crypto/internal/fips140deps/byteorder   [no test files]
?       crypto/internal/fips140deps/cpu [no test files]
?       crypto/internal/fips140deps/godebug     [no test files]
?       crypto/internal/fips140hash     [no test files]
?       crypto/internal/fips140only     [no test files]
?       crypto/internal/impl    [no test files]
?       crypto/internal/randutil        [no test files]
?       crypto/internal/sysrand/internal/seccomp        [no test files]
?       crypto/tls/internal/fips140tls  [no test files]
?       crypto/x509/pkix        [no test files]

##### GOFIPS140=latest go test crypto/...
ok      crypto  2.863s
ok      crypto/aes      0.038s
ok      crypto/cipher   10.242s
ok      crypto/des      0.051s
ok      crypto/dsa      0.017s
ok      crypto/ecdh     0.205s
ok      crypto/ecdsa    0.190s
ok      crypto/ed25519  0.440s
ok      crypto/elliptic 0.040s
?       crypto/fips140  [no test files]
ok      crypto/hkdf     0.045s
ok      crypto/hmac     0.049s
ok      crypto/internal/boring  0.057s
?       crypto/internal/boring/bbig     [no test files]
ok      crypto/internal/boring/bcache   0.885s
?       crypto/internal/boring/sig      [no test files]
?       crypto/internal/cryptotest      [no test files]
?       crypto/internal/entropy [no test files]
?       crypto/internal/fips140 [no test files]
ok      crypto/internal/fips140/aes     0.063s
ok      crypto/internal/fips140/aes/gcm 0.022s [no tests to run]
?       crypto/internal/fips140/alias   [no test files]
ok      crypto/internal/fips140/bigmod  0.231s
?       crypto/internal/fips140/check   [no test files]
?       crypto/internal/fips140/check/checktest [no test files]
ok      crypto/internal/fips140/drbg    0.022s [no tests to run]
ok      crypto/internal/fips140/ecdh    0.022s
ok      crypto/internal/fips140/ecdsa   0.108s
?       crypto/internal/fips140/ed25519 [no test files]
ok      crypto/internal/fips140/edwards25519    0.196s
ok      crypto/internal/fips140/edwards25519/field      0.056s
?       crypto/internal/fips140/hkdf    [no test files]
?       crypto/internal/fips140/hmac    [no test files]
ok      crypto/internal/fips140/mlkem   0.719s
ok      crypto/internal/fips140/nistec  0.020s
ok      crypto/internal/fips140/nistec/fiat     0.015s [no tests to run]
?       crypto/internal/fips140/pbkdf2  [no test files]
ok      crypto/internal/fips140/rsa     0.312s
?       crypto/internal/fips140/sha256  [no test files]
?       crypto/internal/fips140/sha3    [no test files]
?       crypto/internal/fips140/sha512  [no test files]
?       crypto/internal/fips140/ssh     [no test files]
ok      crypto/internal/fips140/subtle  0.027s
?       crypto/internal/fips140/tls12   [no test files]
?       crypto/internal/fips140/tls13   [no test files]
ok      crypto/internal/fips140deps     0.481s
?       crypto/internal/fips140deps/byteorder   [no test files]
?       crypto/internal/fips140deps/cpu [no test files]
?       crypto/internal/fips140deps/godebug     [no test files]
?       crypto/internal/fips140hash     [no test files]
?       crypto/internal/fips140only     [no test files]
ok      crypto/internal/fips140test     2.387s
ok      crypto/internal/hpke    0.020s
?       crypto/internal/impl    [no test files]
?       crypto/internal/randutil        [no test files]
ok      crypto/internal/sysrand 1.154s
?       crypto/internal/sysrand/internal/seccomp        [no test files]
ok      crypto/md5      0.080s
ok      crypto/mlkem    0.339s
ok      crypto/pbkdf2   0.097s
ok      crypto/rand     0.254s
ok      crypto/rc4      0.126s
ok      crypto/rsa      1.800s
ok      crypto/sha1     0.085s
ok      crypto/sha256   0.029s
ok      crypto/sha3     5.642s
ok      crypto/sha512   0.042s
ok      crypto/subtle   0.732s
ok      crypto/tls      9.982s
?       crypto/tls/internal/fips140tls  [no test files]
ok      crypto/x509     1.519s
?       crypto/x509/pkix        [no test files]

##### GOFIPS140=v1.0.0 go test crypto/... # (build and vet only)
?       crypto/fips140  [no test files]
?       crypto/internal/boring/bbig     [no test files]
?       crypto/internal/boring/sig      [no test files]
?       crypto/internal/cryptotest      [no test files]
?       crypto/internal/entropy [no test files]
?       crypto/internal/fips140/v1.0.0  [no test files]
?       crypto/internal/fips140/v1.0.0/alias    [no test files]
?       crypto/internal/fips140/v1.0.0/check    [no test files]
?       crypto/internal/fips140/v1.0.0/check/checktest  [no test files]
?       crypto/internal/fips140/v1.0.0/ed25519  [no test files]
?       crypto/internal/fips140/v1.0.0/hkdf     [no test files]
?       crypto/internal/fips140/v1.0.0/hmac     [no test files]
?       crypto/internal/fips140/v1.0.0/pbkdf2   [no test files]
?       crypto/internal/fips140/v1.0.0/sha256   [no test files]
?       crypto/internal/fips140/v1.0.0/sha3     [no test files]
?       crypto/internal/fips140/v1.0.0/sha512   [no test files]
?       crypto/internal/fips140/v1.0.0/ssh      [no test files]
?       crypto/internal/fips140/v1.0.0/subtle   [no test files]
?       crypto/internal/fips140/v1.0.0/tls12    [no test files]
?       crypto/internal/fips140/v1.0.0/tls13    [no test files]
?       crypto/internal/fips140deps/byteorder   [no test files]
?       crypto/internal/fips140deps/cpu [no test files]
?       crypto/internal/fips140deps/godebug     [no test files]
?       crypto/internal/fips140hash     [no test files]
?       crypto/internal/fips140only     [no test files]
?       crypto/internal/impl    [no test files]
?       crypto/internal/randutil        [no test files]
?       crypto/internal/sysrand/internal/seccomp        [no test files]
?       crypto/tls/internal/fips140tls  [no test files]
?       crypto/x509/pkix        [no test files]

##### Testing without libgcc.
ok      net     0.045s
ok      os/user 0.046s

##### internal linking, -buildmode=pie
ok      reflect 0.107s
ok      crypto/internal/fips140test     0.003s [no tests to run]
ok      os/user 0.032s

##### external linking, -buildmode=exe
ok      crypto/internal/fips140test     0.054s [no tests to run]

##### external linking, -buildmode=pie
ok      crypto/internal/fips140test     0.005s [no tests to run]

##### sync -cpu=10
ok      sync    0.500s

##### Testing cgo
ok      cmd/cgo/internal/test   1.532s
ok      cmd/cgo/internal/test   0.910s
ok      cmd/cgo/internal/testtls        0.018s
ok      cmd/cgo/internal/testtls        0.003s
ok      cmd/cgo/internal/testnocgo      0.010s
ok      cmd/cgo/internal/testnocgo      0.012s
ok      cmd/cgo/internal/test   1.496s
ok      cmd/cgo/internal/test   0.763s
ok      cmd/cgo/internal/test   0.969s
ok      cmd/cgo/internal/test   1.168s
ok      cmd/cgo/internal/testtls        0.009s
ok      cmd/cgo/internal/testnocgo      0.011s

##### API release note check
ok      cmd/relnote     0.060s

##### API check
ok      cmd/api 63.984s

##### GOMAXPROCS=2 runtime -cpu=1 -quick
ok      runtime 34.274s

##### GOMAXPROCS=2 runtime -cpu=2 -quick
ok      runtime 32.342s

##### GOMAXPROCS=2 runtime -cpu=4 -quick
ok      runtime 34.179s

##### Testing race detector
ok      runtime/race    13.151s
ok      flag    1.136s
ok      net     1.099s
ok      os      1.212s
ok      os/exec 2.091s
ok      encoding/gob    1.186s
ok      flag    1.066s
ok      os/exec 2.033s

##### ../test
ok      cmd/internal/testdir    234.551s

ALL TESTS PASSED
---
Installed Go for linux/amd64 in /workspaces/shim-go/src/go
Installed commands in /workspaces/shim-go/src/go/bin
*** You need to add /workspaces/shim-go/src/go/bin to your PATH.
```
