<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace GPBMetadata;

class Graphik
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\Struct::initOnce();
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        \GPBMetadata\Google\Protobuf\Any::initOnce();
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        \GPBMetadata\GithubCom\Mwitkow\GoProtoValidators\Validator::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0ae3380a0d6772617068696b2e70726f746f12036170691a1f676f6f676c" .
            "652f70726f746f6275662f74696d657374616d702e70726f746f1a19676f" .
            "6f676c652f70726f746f6275662f616e792e70726f746f1a1b676f6f676c" .
            "652f70726f746f6275662f656d7074792e70726f746f1a36676974687562" .
            "2e636f6d2f6d7769746b6f772f676f2d70726f746f2d76616c696461746f" .
            "72732f76616c696461746f722e70726f746f22450a03526566121f0a0567" .
            "747970651801200128094210e2df1f0c0a0a5e2e7b312c3232357d24121d" .
            "0a036769641802200128094210e2df1f0c0a0a5e2e7b312c3232357d2422" .
            "3e0a0e526566436f6e7374727563746f72121f0a05677479706518012001" .
            "28094210e2df1f0c0a0a5e2e7b312c3232357d24120b0a03676964180220" .
            "012809221e0a045265667312160a047265667318012003280b32082e6170" .
            "692e52656622510a03446f63121d0a0372656618012001280b32082e6170" .
            "692e5265664206e2df1f022001122b0a0a61747472696275746573180220" .
            "01280b32172e676f6f676c652e70726f746f6275662e5374727563742267" .
            "0a0e446f63436f6e7374727563746f7212280a0372656618012001280b32" .
            "132e6170692e526566436f6e7374727563746f724206e2df1f022001122b" .
            "0a0a6174747269627574657318022001280b32172e676f6f676c652e7072" .
            "6f746f6275662e53747275637422340a0f446f63436f6e7374727563746f" .
            "727312210a04646f637318012003280b32132e6170692e446f63436f6e73" .
            "74727563746f7222610a0954726176657273616c12150a03646f63180120" .
            "01280b32082e6170692e446f6312200a0e74726176657273616c5f706174" .
            "6818022003280b32082e6170692e526566120d0a05646570746818032001" .
            "2804120c0a04686f707318042001280422300a0a54726176657273616c73" .
            "12220a0a74726176657273616c7318012003280b320e2e6170692e547261" .
            "76657273616c22310a04446f637312160a04646f637318012003280b3208" .
            "2e6170692e446f6312110a097365656b5f6e65787418022001280922a801" .
            "0a0a436f6e6e656374696f6e121d0a0372656618012001280b32082e6170" .
            "692e5265664206e2df1f022001122b0a0a61747472696275746573180220" .
            "01280b32172e676f6f676c652e70726f746f6275662e5374727563741210" .
            "0a086469726563746564180320012808121e0a0466726f6d18042001280b" .
            "32082e6170692e5265664206e2df1f022001121c0a02746f18052001280b" .
            "32082e6170692e5265664206e2df1f02200122be010a15436f6e6e656374" .
            "696f6e436f6e7374727563746f7212280a0372656618012001280b32132e" .
            "6170692e526566436f6e7374727563746f724206e2df1f022001122b0a0a" .
            "6174747269627574657318032001280b32172e676f6f676c652e70726f74" .
            "6f6275662e53747275637412100a08646972656374656418042001280812" .
            "1e0a0466726f6d18052001280b32082e6170692e5265664206e2df1f0220" .
            "01121c0a02746f18062001280b32082e6170692e5265664206e2df1f0220" .
            "0122a0010a13536561726368436f6e6e65637446696c746572121b0a0666" .
            "696c74657218012001280b320b2e6170692e46696c746572120d0a056774" .
            "797065180220012809122b0a0a6174747269627574657318032001280b32" .
            "172e676f6f676c652e70726f746f6275662e53747275637412100a086469" .
            "726563746564180420012808121e0a0466726f6d18052001280b32082e61" .
            "70692e5265664206e2df1f0220012282010a15536561726368436f6e6e65" .
            "63744d6546696c746572121b0a0666696c74657218012001280b320b2e61" .
            "70692e46696c746572120d0a056774797065180220012809122b0a0a6174" .
            "747269627574657318032001280b32172e676f6f676c652e70726f746f62" .
            "75662e53747275637412100a08646972656374656418042001280822490a" .
            "16436f6e6e656374696f6e436f6e7374727563746f7273122f0a0b636f6e" .
            "6e656374696f6e7318012003280b321a2e6170692e436f6e6e656374696f" .
            "6e436f6e7374727563746f7222460a0b436f6e6e656374696f6e7312240a" .
            "0b636f6e6e656374696f6e7318012003280b320f2e6170692e436f6e6e65" .
            "6374696f6e12110a097365656b5f6e65787418022001280922e2010a0d43" .
            "6f6e6e65637446696c74657212210a07646f635f72656618012001280b32" .
            "082e6170692e5265664206e2df1f022001121f0a05677479706518022001" .
            "28094210e2df1f0c0a0a5e2e7b312c3232357d2412120a0a657870726573" .
            "73696f6e18032001280912150a056c696d69741804200128044206e2df1f" .
            "02100012430a04736f72741805200128094235e2df1f310a2f28285e7c2c" .
            "2029287c7265662e6769647c7265662e67747970657c5e61747472696275" .
            "7465732e282e2a2929292b24120c0a047365656b180620012809120f0a07" .
            "7265766572736518072001280822c7010a0646696c746572121f0a056774" .
            "7970651801200128094210e2df1f0c0a0a5e2e7b312c3232357d2412120a" .
            "0a65787072657373696f6e18022001280912150a056c696d697418032001" .
            "28044206e2df1f02100012430a04736f72741804200128094235e2df1f31" .
            "0a2f28285e7c2c2029287c7265662e6769647c7265662e67747970657c5e" .
            "617474726962757465732e282e2a2929292b24120c0a047365656b180520" .
            "012809120f0a0772657665727365180620012808120d0a05696e64657818" .
            "07200128092287010a0941676746696c74657212230a0666696c74657218" .
            "012001280b320b2e6170692e46696c7465724206e2df1f02200112210a09" .
            "61676772656761746518022001280e320e2e6170692e4167677265676174" .
            "6512320a056669656c641803200128094223e2df1f1f0a1d28285e7c2c20" .
            "29287c5e617474726962757465732e282e2a2929292b2422ac020a0e5472" .
            "61766572736546696c746572121e0a04726f6f7418012001280b32082e61" .
            "70692e5265664206e2df1f02200112160a0e646f635f6578707265737369" .
            "6f6e180220012809121d0a15636f6e6e656374696f6e5f65787072657373" .
            "696f6e18032001280912150a056c696d69741804200128044206e2df1f02" .
            "100012430a04736f72741805200128094235e2df1f310a2f28285e7c2c20" .
            "29287c7265662e6769647c7265662e67747970657c5e6174747269627574" .
            "65732e282e2a2929292b24120f0a07726576657273651806200128081221" .
            "0a09616c676f726974686d18072001280e320e2e6170692e416c676f7269" .
            "74686d12190a096d61785f64657074681808200128044206e2df1f021000" .
            "12180a086d61785f686f70731809200128044206e2df1f021000228e020a" .
            "1054726176657273654d6546696c74657212160a0e646f635f6578707265" .
            "7373696f6e180120012809121d0a15636f6e6e656374696f6e5f65787072" .
            "657373696f6e18022001280912150a056c696d69741803200128044206e2" .
            "df1f02100012430a04736f72741804200128094235e2df1f310a2f28285e" .
            "7c2c2029287c7265662e6769647c7265662e67747970657c5e6174747269" .
            "62757465732e282e2a2929292b24120f0a07726576657273651805200128" .
            "0812210a09616c676f726974686d18062001280e320e2e6170692e416c67" .
            "6f726974686d12190a096d61785f64657074681807200128044206e2df1f" .
            "02100012180a086d61785f686f70731808200128044206e2df1f02100022" .
            "9c010a10496e646578436f6e7374727563746f72121e0a046e616d651801" .
            "200128094210e2df1f0c0a0a5e2e7b312c3232357d24121f0a0567747970" .
            "651803200128094210e2df1f0c0a0a5e2e7b312c3232357d2412240a0a65" .
            "787072657373696f6e1804200128094210e2df1f0c0a0a5e2e7b312c3232" .
            "357d24120c0a04646f637318062001280812130a0b636f6e6e656374696f" .
            "6e7318072001280822550a0a41757468546172676574121e0a0475736572" .
            "18042001280b32082e6170692e446f634206e2df1f02200112270a067461" .
            "7267657418052001280b32172e676f6f676c652e70726f746f6275662e53" .
            "747275637422a7010a0a417574686f72697a6572121e0a046e616d651801" .
            "200128094210e2df1f0c0a0a5e2e7b312c3232357d2412200a066d657468" .
            "6f641802200128094210e2df1f0c0a0a5e2e7b312c3232357d2412240a0a" .
            "65787072657373696f6e1803200128094210e2df1f0c0a0a5e2e7b312c32" .
            "32357d2412170a0f7461726765745f726571756573747318042001280812" .
            "180a107461726765745f726573706f6e73657318052001280822330a0b41" .
            "7574686f72697a65727312240a0b617574686f72697a6572731801200328" .
            "0b320f2e6170692e417574686f72697a65722299010a0d5479706556616c" .
            "696461746f72121e0a046e616d651801200128094210e2df1f0c0a0a5e2e" .
            "7b312c3232357d24121f0a0567747970651802200128094210e2df1f0c0a" .
            "0a5e2e7b312c3232357d2412240a0a65787072657373696f6e1803200128" .
            "094210e2df1f0c0a0a5e2e7b312c3232357d24120c0a04646f6373180420" .
            "01280812130a0b636f6e6e656374696f6e7318052001280822380a0e5479" .
            "706556616c696461746f727312260a0a76616c696461746f727318012003" .
            "280b32122e6170692e5479706556616c696461746f722291010a05496e64" .
            "6578121e0a046e616d651801200128094210e2df1f0c0a0a5e2e7b312c32" .
            "32357d24121f0a0567747970651803200128094210e2df1f0c0a0a5e2e7b" .
            "312c3232357d2412240a0a65787072657373696f6e1804200128094210e2" .
            "df1f0c0a0a5e2e7b312c3232357d24120c0a04646f637318062001280812" .
            "130a0b636f6e6e656374696f6e7318072001280822260a07496e64657865" .
            "73121b0a07696e646578657318012003280b320a2e6170692e496e646578" .
            "22450a0c53747265616d46696c74657212210a076368616e6e656c180120" .
            "0128094210e2df1f0c0a0a5e2e7b312c3232357d2412120a0a6578707265" .
            "7373696f6e18022001280922470a05477261706812170a04646f63731801" .
            "2001280b32092e6170692e446f637312250a0b636f6e6e656374696f6e73" .
            "18022001280b32102e6170692e436f6e6e656374696f6e7322ed020a0546" .
            "6c61677312190a116f70656e5f69645f646973636f766572791801200128" .
            "0912140a0c73746f726167655f70617468180220012809120f0a076d6574" .
            "7269637318032001280812150a0d616c6c6f775f68656164657273180520" .
            "03280912150a0d616c6c6f775f6d6574686f647318062003280912150a0d" .
            "616c6c6f775f6f726967696e7318072003280912120a0a726f6f745f7573" .
            "65727318082003280912100a08746c735f63657274180920012809120f0a" .
            "07746c735f6b6579180a20012809121c0a14706c617967726f756e645f63" .
            "6c69656e745f6964180b2001280912200a18706c617967726f756e645f63" .
            "6c69656e745f736563726574180c20012809121b0a13706c617967726f75" .
            "6e645f7265646972656374180d2001280912230a1b726571756972655f72" .
            "6571756573745f617574686f72697a657273180f2001280812240a1c7265" .
            "71756972655f726573706f6e73655f617574686f72697a65727318102001" .
            "280822180a07426f6f6c65616e120d0a0576616c75651801200128082217" .
            "0a064e756d626572120d0a0576616c75651801200128012283010a0c4578" .
            "6973747346696c746572121f0a0567747970651801200128094210e2df1f" .
            "0c0a0a5e2e7b312c3232357d2412240a0a65787072657373696f6e180220" .
            "0128094210e2df1f0c0a0a5e2e7b312c3232357d24120c0a047365656b18" .
            "0320012809120f0a0772657665727365180420012808120d0a05696e6465" .
            "7818052001280922520a0445646974121d0a0372656618012001280b3208" .
            "2e6170692e5265664206e2df1f022001122b0a0a61747472696275746573" .
            "18022001280b32172e676f6f676c652e70726f746f6275662e5374727563" .
            "7422560a0a4564697446696c746572121b0a0666696c7465721801200128" .
            "0b320b2e6170692e46696c746572122b0a0a617474726962757465731802" .
            "2001280b32172e676f6f676c652e70726f746f6275662e53747275637422" .
            "170a04506f6e67120f0a076d65737361676518012001280922630a0f4f75" .
            "74626f756e644d65737361676512210a076368616e6e656c180120012809" .
            "4210e2df1f0c0a0a5e2e7b312c3232357d24122d0a046461746118022001" .
            "280b32172e676f6f676c652e70726f746f6275662e5374727563744206e2" .
            "df1f02200122d4010a074d65737361676512210a076368616e6e656c1801" .
            "200128094210e2df1f0c0a0a5e2e7b312c3232357d24122d0a0464617461" .
            "18022001280b32172e676f6f676c652e70726f746f6275662e5374727563" .
            "744206e2df1f022001121e0a047573657218032001280b32082e6170692e" .
            "5265664206e2df1f02200112350a0974696d657374616d7018042001280b" .
            "321a2e676f6f676c652e70726f746f6275662e54696d657374616d704206" .
            "e2df1f02200112200a066d6574686f641805200128094210e2df1f0c0a0a" .
            "5e2e7b312c3232357d2422a4010a06536368656d6112180a10636f6e6e65" .
            "6374696f6e5f747970657318012003280912110a09646f635f7479706573" .
            "18022003280912250a0b617574686f72697a65727318032001280b32102e" .
            "6170692e417574686f72697a65727312270a0a76616c696461746f727318" .
            "042001280b32132e6170692e5479706556616c696461746f7273121d0a07" .
            "696e646578657318052001280b320c2e6170692e496e646578657322200a" .
            "0a4578707246696c74657212120a0a65787072657373696f6e1801200128" .
            "092a1d0a09416c676f726974686d12070a03424653100012070a03444653" .
            "10012a440a0941676772656761746512090a05434f554e54100012070a03" .
            "53554d100112070a03415647100212070a034d4158100312070a034d494e" .
            "100412080a0450524f44100532da100a0f44617461626173655365727669" .
            "6365122b0a0450696e6712162e676f6f676c652e70726f746f6275662e45" .
            "6d7074791a092e6170692e506f6e67220012320a09476574536368656d61" .
            "12162e676f6f676c652e70726f746f6275662e456d7074791a0b2e617069" .
            "2e536368656d612200123c0a0e536574417574686f72697a65727312102e" .
            "6170692e417574686f72697a6572731a162e676f6f676c652e70726f746f" .
            "6275662e456d707479220012340a0a536574496e6465786573120c2e6170" .
            "692e496e64657865731a162e676f6f676c652e70726f746f6275662e456d" .
            "707479220012420a115365745479706556616c696461746f727312132e61" .
            "70692e5479706556616c696461746f72731a162e676f6f676c652e70726f" .
            "746f6275662e456d707479220012280a024d6512162e676f6f676c652e70" .
            "726f746f6275662e456d7074791a082e6170692e446f632200122c0a0943" .
            "7265617465446f6312132e6170692e446f63436f6e7374727563746f721a" .
            "082e6170692e446f632200122f0a0a437265617465446f637312142e6170" .
            "692e446f63436f6e7374727563746f72731a092e6170692e446f63732200" .
            "121e0a06476574446f6312082e6170692e5265661a082e6170692e446f63" .
            "220012260a0a536561726368446f6373120b2e6170692e46696c7465721a" .
            "092e6170692e446f6373220012320a08547261766572736512132e617069" .
            "2e547261766572736546696c7465721a0f2e6170692e5472617665727361" .
            "6c73220012360a0a54726176657273654d6512152e6170692e5472617665" .
            "7273654d6546696c7465721a0f2e6170692e54726176657273616c732200" .
            "12200a0745646974446f6312092e6170692e456469741a082e6170692e44" .
            "6f63220012280a0845646974446f6373120f2e6170692e4564697446696c" .
            "7465721a092e6170692e446f63732200122c0a0644656c446f6312082e61" .
            "70692e5265661a162e676f6f676c652e70726f746f6275662e456d707479" .
            "220012300a0744656c446f6373120b2e6170692e46696c7465721a162e67" .
            "6f6f676c652e70726f746f6275662e456d7074792200122e0a0945786973" .
            "7473446f6312112e6170692e45786973747346696c7465721a0c2e617069" .
            "2e426f6f6c65616e220012350a10457869737473436f6e6e656374696f6e" .
            "12112e6170692e45786973747346696c7465721a0c2e6170692e426f6f6c" .
            "65616e220012220a06486173446f6312082e6170692e5265661a0c2e6170" .
            "692e426f6f6c65616e220012290a0d486173436f6e6e656374696f6e1208" .
            "2e6170692e5265661a0c2e6170692e426f6f6c65616e220012410a104372" .
            "65617465436f6e6e656374696f6e121a2e6170692e436f6e6e656374696f" .
            "6e436f6e7374727563746f721a0f2e6170692e436f6e6e656374696f6e22" .
            "0012440a11437265617465436f6e6e656374696f6e73121b2e6170692e43" .
            "6f6e6e656374696f6e436f6e7374727563746f72731a102e6170692e436f" .
            "6e6e656374696f6e73220012400a10536561726368416e64436f6e6e6563" .
            "7412182e6170692e536561726368436f6e6e65637446696c7465721a102e" .
            "6170692e436f6e6e656374696f6e73220012440a12536561726368416e64" .
            "436f6e6e6563744d65121a2e6170692e536561726368436f6e6e6563744d" .
            "6546696c7465721a102e6170692e436f6e6e656374696f6e732200122c0a" .
            "0d476574436f6e6e656374696f6e12082e6170692e5265661a0f2e617069" .
            "2e436f6e6e656374696f6e220012340a11536561726368436f6e6e656374" .
            "696f6e73120b2e6170692e46696c7465721a102e6170692e436f6e6e6563" .
            "74696f6e732200122e0a0e45646974436f6e6e656374696f6e12092e6170" .
            "692e456469741a0f2e6170692e436f6e6e656374696f6e220012360a0f45" .
            "646974436f6e6e656374696f6e73120f2e6170692e4564697446696c7465" .
            "721a102e6170692e436f6e6e656374696f6e73220012330a0d44656c436f" .
            "6e6e656374696f6e12082e6170692e5265661a162e676f6f676c652e7072" .
            "6f746f6275662e456d707479220012370a0e44656c436f6e6e656374696f" .
            "6e73120b2e6170692e46696c7465721a162e676f6f676c652e70726f746f" .
            "6275662e456d707479220012390a0f436f6e6e656374696f6e7346726f6d" .
            "12122e6170692e436f6e6e65637446696c7465721a102e6170692e436f6e" .
            "6e656374696f6e73220012370a0d436f6e6e656374696f6e73546f12122e" .
            "6170692e436f6e6e65637446696c7465721a102e6170692e436f6e6e6563" .
            "74696f6e732200122e0a0d416767726567617465446f6373120e2e617069" .
            "2e41676746696c7465721a0b2e6170692e4e756d626572220012350a1441" .
            "6767726567617465436f6e6e656374696f6e73120e2e6170692e41676746" .
            "696c7465721a0b2e6170692e4e756d6265722200123b0a0942726f616463" .
            "61737412142e6170692e4f7574626f756e644d6573736167651a162e676f" .
            "6f676c652e70726f746f6275662e456d7074792200122d0a065374726561" .
            "6d12112e6170692e53747265616d46696c7465721a0c2e6170692e4d6573" .
            "7361676522003001123a0a1350757368446f63436f6e7374727563746f72" .
            "7312132e6170692e446f63436f6e7374727563746f721a082e6170692e44" .
            "6f63220028013001124f0a1a50757368436f6e6e656374696f6e436f6e73" .
            "74727563746f7273121a2e6170692e436f6e6e656374696f6e436f6e7374" .
            "727563746f721a0f2e6170692e436f6e6e656374696f6e22002801300112" .
            "300a0853656564446f637312082e6170692e446f631a162e676f6f676c65" .
            "2e70726f746f6275662e456d70747922002801123e0a0f53656564436f6e" .
            "6e656374696f6e73120f2e6170692e436f6e6e656374696f6e1a162e676f" .
            "6f676c652e70726f746f6275662e456d7074792200280142075a05617069" .
            "7062620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}

